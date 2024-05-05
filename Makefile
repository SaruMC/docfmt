TAG=godoc

all: build

.PHONY: build
build:
	docker build -f .docker/Dockerfile -t ${TAG} .

run:
	docker run -it ${TAG}