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

test:
	go test ./... -v

deploy:
	# go get -v github.com/inconshreveable/mousetrap
	- curl -sL https://git.io/goreleaser | rvm 2.4.1 do bash -s -- --release-notes=CHANGELOG.md
	# - curl -X PUT -T devopdays-cli_${VERSION}_linux-i386.deb -umattstratton:${BTKEY} 'https://api.bintray.com/content/devopsdays/deb/devopsdays-cli/${VERSION}/pool/main/d/devopsdays/devopdays-cli_${VERSION}_linux-386.deb;deb_distribution=devopsdays;deb_component=main;deb_architecture=i386;publish=1'
	# - curl -X PUT -T devopdays-cli_${VERSION}_linux-amd64.deb -umattstratton:${BTKEY} 'https://api.bintray.com/content/devopsdays/deb/devopsdays-cli/${VERSION}/pool/main/d/devopsdays/devopdays-cli_${VERSION}_linux-amd64.deb;deb_distribution=devopsdays;deb_component=main;deb_architecture=amd64;publish=1'
	# - curl PUT -T devopdays-cli_$VERSION_linux-i386.rpm -umattstratton:{BTKEY} 'https://api.bintray.com/content/devopsdays/rpm/devopsdays-cli/${VERSION}/devopdays-cli_$VERSION_linux-386.rpm;publish=1'
	# - curl PUT -T devopdays-cli_$VERSION_linux-amd64.rpm -umattstratton:{BTKEY} 'https://api.bintray.com/content/devopsdays/rpm/devopsdays-cli/${VERSION}/devopdays-cli_$VERSION_linux-amd64.rpm;publish=1'

travis:
	$(HOME)/gopath/bin/goveralls -service=travis-ci
