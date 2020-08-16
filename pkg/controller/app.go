package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"sync"

	"github.com/rahulsidpatil/qlikapp/pkg/dal"
	"github.com/rahulsidpatil/qlikapp/pkg/util"
	_ "github.com/rahulsidpatil/qlikapp/pkg/util"

	"github.com/gorilla/mux"
	"github.com/rahulsidpatil/qlikapp/api/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

var swaggerAddr, svcHost, svcAddr, svcPort, metricsAddr, metricsPort, svcPathPrefix, svcVersion string

func init() {
	svcVersion = os.Getenv("SVC_VERSION")
	svcHost = os.Getenv("SVC_HOST")
	svcPort = os.Getenv("SVC_PORT")
	metricsPort = os.Getenv("METRICS_PORT")
	svcAddr = svcHost + ":" + svcPort
	metricsAddr = svcHost + ":" + metricsPort
	svcPathPrefix = svcVersion + "/" + os.Getenv("SVC_PATH_PREFIX")
	swaggerAddr = "http://localhost:" + svcPort + "/swagger/doc.json"
}

type App struct {
	Router *mux.Router
	DB     dal.Interface
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.DB = dal.GetMySQLDriver()
	a.setSwaggerInfo()
	a.initializeRoutes()
}

func (a *App) Run() {
	wg := sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		if i == 0 {
			log.Println("Message Service is being started at:", svcAddr)
			go func() {
				defer wg.Done()
				log.Println(http.ListenAndServe(svcAddr, a.Router))
			}()
		} else {
			log.Println("Metrics Service is being started at:", metricsAddr)
			go func() {
				defer wg.Done()
				log.Println(http.ListenAndServe(metricsAddr, nil))
			}()
		}
	}
	wg.Wait()
}

func (a *App) initializeRoutes() {

	a.Router.HandleFunc(os.Getenv("SVC_VERSION")+"/hello", a.hello).Methods("GET")
	a.Router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL(swaggerAddr), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))
	a.Router.HandleFunc(svcPathPrefix, a.getAll).Methods("GET")
	a.Router.HandleFunc(svcPathPrefix, a.addMessage).Methods("POST")
	a.Router.HandleFunc(svcPathPrefix+"/{id:[0-9]+}", a.getMessage).Methods("GET")
	a.Router.HandleFunc(svcPathPrefix+"/{id:[0-9]+}", a.updateMessage).Methods("PUT")
	a.Router.HandleFunc(svcPathPrefix+"/{id:[0-9]+}", a.deleteMessage).Methods("DELETE")
	a.Router.HandleFunc(svcPathPrefix+"/palindromeChk/{id:[0-9]+}", a.palindromeChk).Methods("GET")
}

func (a *App) setSwaggerInfo() {
	// programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample qlikapp server."
	docs.SwaggerInfo.Version = "1.0"
	//TODO: remove hard-coding of Host address
	docs.SwaggerInfo.Host = "localhost:" + svcPort
	docs.SwaggerInfo.BasePath = svcVersion
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

// @Summary Say hello to user
// @Description Say hello to user
// @Success 200 {object} string
// @Failure 404
// @Router /hello [get]
func (a *App) hello(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, "Hello qlikapp..!!!")
}

// @Summary Get all messages
// @Description Get all messages
// @Success 200 {array} dal.Message
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /messages [get]
func (a *App) getAll(w http.ResponseWriter, r *http.Request) {
	messages, err := a.DB.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, messages)
}

// @Summary Add new messages
// @Description Add new messages
// @Accept  json
// @Produce  json
// @Param message body dal.Message true "Add message"
// @Success 200 {object} dal.Message
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /messages [post]
func (a *App) addMessage(w http.ResponseWriter, r *http.Request) {
	var msg dal.Message
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&msg); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := a.DB.AddMessage(&msg); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, msg)
}

// @Summary Fetch message by ID
// @Description Fetch message by ID
// @Accept  json
// @Produce  json
// @Param id path int true "Message ID"
// @Success 200 {object} dal.Message
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /messages/{id} [get]
func (a *App) getMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid message id")
		return
	}

	msg := dal.Message{ID: id}
	if err := a.DB.GetMessage(&msg); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Message not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, msg)
}

// @Summary Check if the message specified by ID is a palindrome or not
// @Description Check if the message specified by ID is a palindrome or not
// @Accept  json
// @Produce  json
// @Param id path int true "Message ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /messages/palindromeChk/{id} [get]
func (a *App) palindromeChk(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid message id")
		return
	}

	msg := dal.Message{ID: id}
	if err := a.DB.GetMessage(&msg); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Message not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	palindromeChk := util.Palindrome(msg.Message)
	response := map[string]interface{}{
		"Message":    msg,
		"Palindrome": palindromeChk,
	}
	respondWithJSON(w, http.StatusOK, response)
}

// @Summary Update message by ID
// @Description Update message by ID
// @Accept  json
// @Produce  json
// @Param id path int true "Message ID"
// @Param message body dal.Message true "Update message"
// @Success 200 {object} dal.Message
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /messages/{id} [put]
func (a *App) updateMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid message ID")
		return
	}

	var msg dal.Message
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&msg); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	msg.ID = id

	if err := a.DB.UpdateMessage(&msg); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, msg)
}

// @Summary Delete message by ID
// @Description Delete message by ID
// @Accept  json
// @Produce  json
// @Param id path int true "Message ID"
// @Success 200 {object} dal.Message
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /messages/{id} [delete]
func (a *App) deleteMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}

	msg := dal.Message{ID: id}
	if err := a.DB.DeleteMessage(&msg); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
