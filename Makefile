API=chasse-api
APP=chasse-app
BUILD_DIR=build

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
	tar -zcvf ${APP}.tar.gz build;\
	mv ${APP}.tar.gz ../${BUILD_DIR}

build-api: 
	if ! [ -d ${BUILD_DIR} ] ; then mkdir ${BUILD_DIR} ; fi
	cd ${API};\
	go mod tidy;\
	GOOS=linux GOARCH=amd64 go build -o ../build/${API} main.go

build: build-api build-app 

run-app:
	cd ${APP};\
	npm install;\
	REACT_APP_DEV_MODE=true npm run start

run-app-remote:
	cd ${APP};\
	npm install;\
	npm run start

run-api: 
	docker-compose up -d
	cd ${API};\
	if ! [ -d ${BUILD_DIR} ] ; then mkdir ${BUILD_DIR} ; fi;\
	go mod tidy;\
	go build -o build/${API} .;\
	./${BUILD_DIR}/${API}

run: run-api run-app
