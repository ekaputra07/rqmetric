GO111MODULE=on

develop:
	go get -u golang.org/x/lint/golint

check:
	golint
	go vet
	go fmt
	go test

build:
	go build

install:
	go install

linux:
	env GOOS=linux GOARCH=amd64 go build