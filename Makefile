GO111MODULE=on

develop:
	go get -u github.com/gobuffalo/packr/packr

build:
	packr build

install:
	packr install