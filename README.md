# qlikapp: sample message server for QlikAudition
The application is a simple REST API server that will provide endpoints to allow creating, accessing and manipulating ‘messages’. The application also provides an endpoint to observe `application's runtime profiling data`

# Getting started
Following are the steps to run this application in local docker environment

## Prerequisites
- git (recommended version version 2.17.1)
- Go (recommended version version go1.13.10)
- docker (recommended version 19.03.12). 
Installation steps available at: https://docs.docker.com/engine/install/ubuntu/ https://docs.docker.com/engine/install/linux-postinstall/
Make sure to be able to run `docker as non root user` 
- docker-compose version (recommended version 1.26.2)
Installation steps available at: https://docs.docker.com/compose/install/ https://docs.docker.com/engine/security/rootless/
Make sure to be able to run `docker-compose as non root user`

## Building, Running and Accessing the application:
1) Clone the qlikapp repository to your $(GOPATH) e.g. `~/go/src/github.com/rahulsidpatil`
```
git clone git@github.com:rahulsidpatil/qlikapp.git
```
2) cd to qlikapp dir
```
cd qlikapp
```
3) Run follwoing command
```
make docker-deploy-up
```
It will a take a while to build and deploy the application on your local docker environment.
Once the deployment is successfull; the console will display application access urls as:
```
echo "Server started at url: http://localhost:8080"
echo "The API documentation is available at url: http://localhost:8080/swagger/"
echo "Server runtime profiling data available at url: http://localhost:6060/debug/pprof"

```
## qlikapp in action:


# Licensing
qlikapp is under the MIT License.
