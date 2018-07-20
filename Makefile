deps:
	go get github.com/golang/mock/mockgen/model
	go install github.com/golang/mock/mockgen

mocks:
	mockgen -package mocks github.com/aws/aws-sdk-go/service/s3/s3iface S3API > mocks/mock_s3.go

.PHONY: deps mocks
