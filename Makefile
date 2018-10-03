DOCKER_BUILD_TAG=ledmatrix-build

all: build compile

build:
	docker rmi -f ${DOCKER_BUILD_TAG}
	docker build --rm -t ${DOCKER_BUILD_TAG} .

compile:
	docker run --rm -v $(PWD):/src ${DOCKER_BUILD_TAG}

clean:
	rm -f bin/*
	docker rmi -f ${DOCKER_BUILD_TAG}
