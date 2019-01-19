GO111MODULE=on

develop:
	go get -u golang.org/x/lint/golint
	go get -u github.com/gobuffalo/packr/packr

check:
	golint
	go vet
	go fmt
	go test

build:
	packr build

install:
	packr install

linux:
	env GOOS=linux GOARCH=amd64 packr build