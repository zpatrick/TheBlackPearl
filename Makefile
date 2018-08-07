VERSION?=$(shell git describe --tags --always)
DOCKER_IMAGE=zpatrick/tbp

deps:
	go get github.com/golang/mock/mockgen/model
	go install github.com/golang/mock/mockgen

mocks:
	mockgen -package mocks github.com/aws/aws-sdk-go/service/s3/s3iface S3API > mocks/mock_s3.go

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a --ldflags "-X main.Version=$(VERSION)" -o tbp . 
	docker build -t $(DOCKER_IMAGE):$(VERSION) .

release: build
	docker push $(DOCKER_IMAGE):$(VERSION)
	docker tag  $(DOCKER_IMAGE):$(VERSION) $(DOCKER_IMAGE):latest
	docker push $(DOCKER_IMAGE):latest

.PHONY: deps mocks build release
