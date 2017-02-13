NAME:=$(shell basename `git rev-parse --show-toplevel`)
BUILDERNAME=alpine-golang
RELEASE:=$(shell git rev-parse --verify --short HEAD)
USER=bketelsen
PWD:=$(shell pwd)


all: docker-push

builder: builder-push

builder-clean:
	docker rmi ${BUILDERNAME} &>/dev/null || true
	docker rmi ${NAME} &>/dev/null || true

builder-build: builder-clean
	docker build --no-cache --pull=true -t ${USER}/${BUILDERNAME}:${RELEASE} . -f Dockerfile.build
	docker tag ${USER}/${BUILDERNAME}:${RELEASE} ${USER}/${BUILDERNAME}:latest

builder-push: builder-build
	docker login -u ${USER}
	docker push ${USER}/${BUILDERNAME}:${RELEASE}
	docker push ${USER}/${BUILDERNAME}:latest


clean:
	rm -rf bin/${NAME}

test: clean
	GO_ENV=test APP_PATH=./ go test ./...

build:
	npm rebuild node-sass
	npm i
	go get ./...
	go get github.com/gobuffalo/buffalo/...
	buffalo build


docker-clean:
	docker rmi ${NAME} &>/dev/null || true

docker-build: docker-clean
	docker run --rm -it -v ${PWD}:/root/src/github.com/gopheracademy/learn -w /root/src/github.com/gopheracademy/learn ${USER}/alpine-golang:latest make build
	docker build --pull=true --no-cache -t ${USER}/${NAME}:${RELEASE} .
	docker tag ${USER}/${NAME}:${RELEASE} ${USER}/${NAME}:latest


docker-push: docker-build
	docker push ${USER}/${NAME}:${RELEASE}
	docker push ${USER}/${NAME}:latest
