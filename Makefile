.PHONY: all alert_manager test

DOCKER_IMAGE := mayuresh82/alert_manager
VERSION := $(shell git describe --exact-match --tags 2>/dev/null)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT := $(shell git rev-parse --short HEAD)
DOCKER_TAG := $(COMMIT)
LDFLAGS := $(LDFLAGS) -X main.commit=$(COMMIT) -X main.branch=$(BRANCH)
ifdef VERSION
    LDFLAGS += -X main.version=$(VERSION)
	DOCKER_TAG = $(VERSION)
endif

all:
	$(MAKE) deps
	$(MAKE) alert_manager

deps:
	go get -u golang.org/x/lint/golint
	go get -u github.com/golang/dep/cmd/dep
	dep ensure
	
docker:
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

alert_manager:
	go build -ldflags "$(LDFLAGS)" ./cmd/alert_manager

debug:
	dep ensure
	go build -race ./cmd/alert_manager

test:
	go test -v -race -short -failfast ./...

linux:
	dep ensure
	GOOS=linux GOARCH=amd64 go build -o alert_manager_linux ./cmd/alert_manager
