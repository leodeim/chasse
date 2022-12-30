include .env

API=chasse-api
APP=chasse-app
BUILD_DIR=build
GIT_SHA_FETCH := $(shell git rev-parse HEAD | cut -c 1-8)
export GIT_SHA=$(GIT_SHA_FETCH)
API_FILE_NAME=${API}-${GIT_SHA}
APP_FILE_NAME=${APP}-${GIT_SHA}.tar.gz
HEALTH=

clean:
	docker-compose down
	if [ -d ${APP}/${BUILD_DIR} ] ; then rm -rf ${APP}/${BUILD_DIR} ; fi
	if [ -d ${API}/${BUILD_DIR} ] ; then rm -rf ${API}/${BUILD_DIR} ; fi
	if [ -d ${BUILD_DIR} ] ; then rm -rf ${BUILD_DIR} ; fi

build-app:
	if ! [ -d ${BUILD_DIR} ] ; then mkdir ${BUILD_DIR} ; fi
	cd ${APP};\
	npm install;\
	npm run build;\
	tar -zcvf ${APP_FILE_NAME} build;\
	mv ${APP_FILE_NAME} ../${BUILD_DIR};

build-api: 
	if ! [ -d ${BUILD_DIR} ] ; then mkdir ${BUILD_DIR} ; fi
	cd ${API};\
	go mod tidy;\
	GOOS=linux GOARCH=amd64 go build -o ../build/${API_FILE_NAME} .;

build: build-api build-app 

run-app:
	cd ${APP};\
	npm install;\
	REACT_APP_DEV_MODE=true npm run start;

run-app-remote:
	cd ${APP};\
	npm install;\
	npm run start;

run-api: 
	docker-compose up -d
	cd ${API};\
	if ! [ -d ${BUILD_DIR} ] ; then mkdir ${BUILD_DIR} ; fi;\
	go mod tidy;\
	go build -o build/${API} .;\
	./${BUILD_DIR}/${API};

run: run-api run-app

deploy-api:
	scp build/${API_FILE_NAME} root@${SERVER_URL}:~/bin
	ssh root@${SERVER_URL} "./deploy-api.sh bin/${API_FILE_NAME}"

deploy-app:
	scp build/${APP_FILE_NAME} root@${SERVER_URL}:~/bin
	ssh root@${SERVER_URL} "./deploy-app.sh bin/${APP_FILE_NAME}"

build-deploy-api: build-api deploy-api

build-deploy-app: build-app deploy-app

deploy: deploy-api deploy-app

build-deploy: build deploy

api-health:
	./misc/scripts/check_health.sh ${GIT_SHA} ${HEALTH_URL}

git_sha:
	@echo ${GIT_SHA}
