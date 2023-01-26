include .env

API=chasse-api
APP=chasse-app
BUILD_DIR=build
GIT_SHA_FETCH := $(shell git rev-parse HEAD | cut -c 1-8)
export GIT_SHA=$(GIT_SHA_FETCH)
API_FILE_NAME=${API}-${GIT_SHA}
APP_FILE_NAME=${APP}-${GIT_SHA}.tar.gz
HEALTH=

.PHONY: clean
clean:
	docker-compose down
	if [ -d ${APP}/${BUILD_DIR} ] ; then rm -rf ${APP}/${BUILD_DIR} ; fi
	if [ -d ${API}/${BUILD_DIR} ] ; then rm -rf ${API}/${BUILD_DIR} ; fi
	if [ -d ${BUILD_DIR} ] ; then rm -rf ${BUILD_DIR} ; fi

.PHONY: build-app
build-app:
	if ! [ -d ${BUILD_DIR} ] ; then mkdir ${BUILD_DIR} ; fi
	cd ${APP};\
	npm install;\
	REACT_APP_VERSION=ver_${GIT_SHA} REACT_APP_API_URL=${REACT_APP_API_URL} REACT_APP_WS_URL=${REACT_APP_WS_URL}  npm run build;\
	tar -zcvf ${APP_FILE_NAME} build;\
	mv ${APP_FILE_NAME} ../${BUILD_DIR};

.PHONY: build-api
build-api: 
	if ! [ -d ${BUILD_DIR} ] ; then mkdir ${BUILD_DIR} ; fi
	cd ${API};\
	go mod tidy;\
	GOOS=linux GOARCH=amd64 go build -o ../build/${API_FILE_NAME} .;

.PHONY: build
build: build-api build-app 

.PHONY: run-app
run-app:
	cd ${APP};\
	npm install;\
	REACT_APP_VERSION=ver_dev_ REACT_APP_DEV_MODE=true npm run start;

.PHONY: run-app-remote
run-app-remote:
	cd ${APP};\
	npm install;\
	REACT_APP_VERSION=ver_dev_ REACT_APP_DEV_MODE=true REACT_APP_API_URL=${REACT_APP_API_URL} REACT_APP_WS_URL=${REACT_APP_WS_URL} npm run start;

.PHONY: run-api
run-api: 
	docker-compose up -d
	cd ${API};\
	if ! [ -d ${BUILD_DIR} ] ; then mkdir ${BUILD_DIR} ; fi;\
	go mod tidy;\
	go build -o build/${API} .;\
	./${BUILD_DIR}/${API};

.PHONY: run
run: run-api run-app

.PHONY: deploy-api
deploy-api:
	scp build/${API_FILE_NAME} root@${SERVER_URL}:~/bin
	ssh root@${SERVER_URL} "./deploy-api.sh bin/${API_FILE_NAME}"
	misc/scripts/deploy-notifier.sh ${AIRBRAKE_ID} ${AIRBRAKE_KEY} api ${GIT_SHA}

.PHONY: deploy-app
deploy-app:
	scp build/${APP_FILE_NAME} root@${SERVER_URL}:~/bin
	ssh root@${SERVER_URL} "./deploy-app.sh bin/${APP_FILE_NAME}"
	misc/scripts/deploy-notifier.sh ${AIRBRAKE_ID} ${AIRBRAKE_KEY} app ${GIT_SHA}

.PHONY: build-deploy-api
build-deploy-api: build-api deploy-api

.PHONY: build-deploy-app
build-deploy-app: build-app deploy-app

.PHONY: deploy
deploy: deploy-api deploy-app

.PHONY: build-deploy
build-deploy: build deploy

.PHONY: health-api
health-api:
	./misc/scripts/check-health-api.sh ${GIT_SHA} ${HEALTH_URL}

.PHONY: health-app
health-app:
	./misc/scripts/check-health-app.sh 200 ${APP_URL}

.PHONY: health
health: health-api health-app
