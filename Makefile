package = github.com/mattstratton/devopdays

BINARY=devopsdays
VERSION=`git describe --tags`
BUILD=`date +%FT%T%z`

LDFLAGS=-ldflags "-w -s -X github.com/mattstratton/devopsdays/cmd.Version=${VERSION} -X github.com/mattstratton/devopsdays/cmd.Build=${BUILD}"

build:
	go build ${LDFLAGS} -o release/${BINARY} $(package)

.PHONY: install release test travis

install:
	go get -t -v ./...

release:
	go get -v github.com/inconshreveable/mousetrap
	rm -rf build/devopsdays
	mkdir -p build/devopsdays
	mkdir -p build/linux-amd64
	mkdir -p build/linux-386
	mkdir -p build/darwin-amd64
	mkdir -p build/darwin-386
	mkdir -p build/windows-amd64
	mkdir -p build/windows-386

	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o build/linux-amd64/devopsdays_${VERSION} $(package)
	GOOS=linux GOARCH=386 go build ${LDFLAGS} -o build/linux-386/devopsdays_${VERSION} $(package)
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o build/darwin-amd64/devopsdays_${VERSION} $(package)
	GOOS=darwin GOARCH=386 go build ${LDFLAGS} -o build/darwin-386/devopsdays_${VERSION} $(package)
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o build/windows-amd64/devopsdays_${VERSION}.exe $(package)
	GOOS=windows GOARCH=386 go build ${LDFLAGS} -o build/windows-386/devopsdays_${VERSION}.exe $(package)

	zip -r build/linux-amd64/devopsdays_${VERSION}.zip release/linux-amd64_devopsdays_${VERSION}
	zip -r build/linux-amd64/devopsdays_${VERSION}.zip release/linux-386_devopsdays_${VERSION}
	zip -r build/linux-amd64/devopsdays_${VERSION}.zip release/darwin-amd64_devopsdays_${VERSION}
	zip -r build/linux-amd64/devopsdays_${VERSION}.zip release/darwin-386_devopsdays_${VERSION}
	zip -r build/linux-amd64/devopsdays_${VERSION}.zip release/windows-amd64_devopsdays_${VERSION}
	zip -r build/linux-amd64/devopsdays_${VERSION}.zip release/windows-386_devopsdays_${VERSION}

	ls release

test:
	go test -v

travis:
	$(HOME)/gopath/bin/goveralls -service=travis-ci
