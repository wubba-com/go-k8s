PROJECT?=github.com/wubba-com/go-k8s
APP?=docker_gs_go
PORT?=4500

RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
CONTAINER_IMAGE?=devise3000/${APP}

GOOS?=linux
GOARCH?=amd64

clean:
	docker rm -f ${APP}

build:
	docker build -t $(CONTAINER_IMAGE):$(RELEASE) .

run: build
	docker stop $(APP):$(RELEASE) || true && docker rm $(APP):$(RELEASE) || true
	docker run --name ${APP} -p 80:${PORT} --rm $(CONTAINER_IMAGE):$(RELEASE)

push: build
	docker push $(CONTAINER_IMAGE):$(RELEASE)

stop:
	docker stop $(APP)

minikube: push
	for t in $(shell find ./kubernetes -type f -name "*.yaml"); do \
        cat $$t | \
        	gsed -E "s/\{\{(\s*)\.Release(\s*)\}\}/$(RELEASE)/g" | \
        	gsed -E "s/\{\{(\s*)\.ServiceName(\s*)\}\}/$(APP)/g"; \
        echo ---; \
    done > tmp.yaml
	kubectl apply -f tmp.yaml