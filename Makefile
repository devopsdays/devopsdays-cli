package = github.com/devopsdays/devopsdays-cli

SOURCE_FILES?=./...
TEST_PATTERN?=.
TEST_OPTIONS?=

BINARY=devopsdays-cli

VERSION:=$(shell git describe --tags)
BUILD:=$(shell date +%FT%T%z)

LDFLAGS=-ldflags "-w -s -X github.com/devopsdays/devopsdays-cli/cmd.Version=${VERSION} -X github.com/devopsdays/devopsdays-cli/cmd.Build=${BUILD}"

# Install all the build and lint dependencies
setup:
	go get -u github.com/alecthomas/gometalinter
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/pierrre/gotestcover
	go get -u golang.org/x/tools/cmd/cover
	go get -u github.com/inconshreveable/mousetrap
	go get -u github.com/mattn/goveralls
	# dep ensure
	gometalinter --install

build:
	go build ${LDFLAGS} -o release/${BINARY} $(package)

.PHONY: install release test travis

install:
	go get -t -v ./...

test:
	gotestcover $(TEST_OPTIONS) -covermode=atomic -coverprofile=coverage.txt $(SOURCE_FILES) -run $(TEST_PATTERN) -timeout=2m

# Run all the tests and opens the coverage report
cover: test
	go tool cover -html=coverage.txt 

# gofmt and goimports all go files
fmt:
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

# Run all the linters
lint:
	gometalinter --vendor ./...

# Run all the tests and code checks
# ci: test lint
ci: test # will add lint later

deploy:
	# go get -v github.com/inconshreveable/mousetrap
	- curl -sL https://git.io/goreleaser | rvm 2.4.1 do bash -s -- --release-notes=RELEASE_NOTES.md
	# - curl -X PUT -T devopdays-cli_${VERSION}_linux-i386.deb -umattstratton:${BTKEY} 'https://api.bintray.com/content/devopsdays/deb/devopsdays-cli/${VERSION}/pool/main/d/devopsdays/devopdays-cli_${VERSION}_linux-386.deb;deb_distribution=devopsdays;deb_component=main;deb_architecture=i386;publish=1'
	# - curl -X PUT -T devopdays-cli_${VERSION}_linux-amd64.deb -umattstratton:${BTKEY} 'https://api.bintray.com/content/devopsdays/deb/devopsdays-cli/${VERSION}/pool/main/d/devopsdays/devopdays-cli_${VERSION}_linux-amd64.deb;deb_distribution=devopsdays;deb_component=main;deb_architecture=amd64;publish=1'
	# - curl PUT -T devopdays-cli_$VERSION_linux-i386.rpm -umattstratton:{BTKEY} 'https://api.bintray.com/content/devopsdays/rpm/devopsdays-cli/${VERSION}/devopdays-cli_$VERSION_linux-386.rpm;publish=1'
	# - curl PUT -T devopdays-cli_$VERSION_linux-amd64.rpm -umattstratton:{BTKEY} 'https://api.bintray.com/content/devopsdays/rpm/devopsdays-cli/${VERSION}/devopdays-cli_$VERSION_linux-amd64.rpm;publish=1'

travis:
	$(HOME)/gopath/bin/goveralls -service=travis-ci
