package = github.com/mattstratton/probablyfine

BINARY=probablyfine
VERSION=`git describe --tags`
BUILD=`date +%FT%T%z`

LDFLAGS=-ldflags "-w -s -X github.com/mattstratton/probablyfine/cmd.Version=${VERSION} -X github.com/mattstratton/probablyfine/cmd.Build=${BUILD}"

build:
	go build ${LDFLAGS} -o release/${BINARY} $(package)

.PHONY: install release test travis

install:
	go get -t -v ./...

release:
	go get -v github.com/inconshreveable/mousetrap
	mkdir -p release
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o release/probablyfine_${VERSION}-linux-amd64 $(package)
	GOOS=linux GOARCH=386 go build ${LDFLAGS} -o release/probablyfine_${VERSION}-linux-386 $(package)
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o release/probablyfine_${VERSION}-darwin-amd64 $(package)
	GOOS=darwin GOARCH=386 go build ${LDFLAGS} -o release/probablyfine_${VERSION}-darwin-386 $(package)
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o release/probablyfine_${VERSION}-windows-amd64.exe $(package)
	GOOS=windows GOARCH=386 go build ${LDFLAGS} -o release/probablyfine_${VERSION}-windows-386.exe $(package)
	cd release
	ls

test:
	go test -v

travis:
	$(HOME)/gopath/bin/goveralls -service=travis-ci
