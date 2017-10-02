package = github.com/devopsdays/devopsdays-cli

BINARY=devopsdays-cli

VERSION:=$(shell git describe --tags)
BUILD:=$(shell date +%FT%T%z)

LDFLAGS=-ldflags "-w -s -X github.com/devopsdays/devopsdays-cli/cmd.Version=${VERSION} -X github.com/devopsdays/devopsdays-cli/cmd.Build=${BUILD}"

build:
	go build ${LDFLAGS} -o release/${BINARY} $(package)

.PHONY: install release test travis

install:
	go get -t -v ./...

release:
	go get -v github.com/inconshreveable/mousetrap
	rm -rf build/devopsdays-cli
	mkdir -p build/devopsdays-cli
	mkdir -p build/linux-amd64
	mkdir -p build/linux-386
	mkdir -p build/darwin-amd64
	mkdir -p build/darwin-386
	mkdir -p build/windows-amd64
	mkdir -p build/windows-386

	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o build/linux-amd64/devopsdays-cli_${VERSION} $(package)
	GOOS=linux GOARCH=386 go build ${LDFLAGS} -o build/linux-386/devopsdays-cli_${VERSION} $(package)
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o build/darwin-amd64/devopsdays-cli_${VERSION} $(package)
	GOOS=darwin GOARCH=386 go build ${LDFLAGS} -o build/darwin-386/devopsdays-cli_${VERSION} $(package)
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o build/windows-amd64/devopsdays-cli_${VERSION}.exe $(package)
	GOOS=windows GOARCH=386 go build ${LDFLAGS} -o build/windows-386/devopsdays-cli_${VERSION}.exe $(package)

	zip -r build/linux-amd64/devopsdays-cli_${VERSION}.zip release/linux-amd64_devopsdays-cli_${VERSION}
	zip -r build/linux-amd64/devopsdays-cli_${VERSION}.zip release/linux-386_devopsdays-cli_${VERSION}
	zip -r build/linux-amd64/devopsdays-cli_${VERSION}.zip release/darwin-amd64_devopsdays-cli_${VERSION}
	zip -r build/linux-amd64/devopsdays-cli_${VERSION}.zip release/darwin-386_devopsdays-cli_${VERSION}
	zip -r build/linux-amd64/devopsdays-cli_${VERSION}.zip release/windows-amd64_devopsdays-cli_${VERSION}
	zip -r build/linux-amd64/devopsdays-cli_${VERSION}.zip release/windows-386_devopsdays-cli_${VERSION}

	ls release

test:
	go test -v

deploy:
	# go get -v github.com/inconshreveable/mousetrap
	- curl -sL https://git.io/goreleaser | rvm 2.4.1 do bash
	- ls dist
	- ./util/jfrog bt pc --key=$BTKEY --user=devopsdays --licenses=MIT --vcs-url=https://github.com/devopsdays/rpm devopsdays/rpm/$GH_APP || echo "package already exists"
	- ./util/jfrog bt upload --override=true --key $BTKEY --publish=true dist/devopdays-cli_$VERSION_linux-386.rpm devopsdays/rpm/devopdays-cli/$VERSION pool/$POOL/devopdays-cli/
	- ./util/jfrog bt upload --override=true --key $BTKEY --publish=true dist/devopdays-cli_$VERSION_linux-amd64.rpm devopsdays/rpm/devopdays-cli/$VERSION pool/$POOL/devopdays-cli/
	- ./util/jfrog bt pc --key=$BTKEY --user=devopsdays --licenses=MIT --vcs-url=https://github.com/devopsdays/deb devopsdays/deb/devopdays-cli || echo "package already exists"
	- ./util/jfrog bt upload --override=true --key $BTKEY --publish=true --deb=unstable/main/386 dist/devopdays-cli_$VERSION_linux-386.deb devopsdays/deb/devopdays-cli/$VERSION pool/$POOL/devopdays-cli/
	- ./util/jfrog bt upload --override=true --key $BTKEY --publish=true --deb=unstable/main/amd64 dist/devopdays-cli_$VERSION_linux-amd64.deb devopsdays/deb/devopdays-cli/$VERSION pool/$POOL/devopdays-cli/
	- curl -X POST -u mattstratton:${BTKEY} https://api.bintray.com/calc_metadata/devopsdays/deb
	- curl -X POST -u mattstratton:${BTKEY} https://api.bintray.com/calc_metadata/devopsdays/rpm
travis:
	$(HOME)/gopath/bin/goveralls -service=travis-ci
