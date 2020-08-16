all: clean swagger-update build

.PHONY: build

swagger-update:
	swag init --parseDependency -d ./cmd/ -o ./api/docs

# ensure the changes are buildable
build:
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build  -o ./qlikapp cmd/main.go

# build images (environment images are not included)
images:
	docker build -t rahulsidpatil/qlikapp:latest -f ./Dockerfile .
	docker build -t rahulsidpatil/qlikdb:latest ./build/db/mysql/.

docker-deploy-up:
	docker-compose -f ./build/docker-deploy/docker-deploy.yaml up --build -d
	echo "Server started at url: http://localhost:8080"
	echo "The API documentation is available at url: http://localhost:8080/swagger/"
	echo "Server runtime profiling data available at url: http://localhost:6060/debug/pprof"

docker-deploy-down:
	docker-compose -f ./build/docker-deploy/docker-deploy.yaml down

clean:
	@rm -f ./qlikapp