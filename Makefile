GO111MODULE=on

develop:
	go get -u github.com/gobuffalo/packr/packr

check:
	go fmt && go test

build:
	packr build

install:
	packr install