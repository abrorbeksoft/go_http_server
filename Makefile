CURRENT_DIR=$(shell pwd)

APP=$(shell basename ${CURRENT_DIR})

APP_CMD_DIR=${CURRENT_DIR}/cmd

PROJECT_NAME=${PROJECT_NAME}
REGISTRY=${REGISTRY}
TAG=latest
ENV_TAG=latest



build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

start:
	  go run cmd/main.go

.PHONY:start