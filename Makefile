package = github.com/devopsdays/devopsdays-cli

PACKAGES=$(shell go list ./... | grep -v /vendor/)
SOURCE_FILES?=./...
TEST_PATTERN?=.
TEST_OPTIONS?=

BINARY=devopsdays-cli

VERSION:=$(shell git describe --tags)
# VERSION := $(shell grep "const Version " version.go | sed -E 's/.*"(.+)"$$/\1/')
BUILD:=$(shell date +%FT%T%z)

LDFLAGS=-ldflags "-w -s -X github.com/devopsdays/devopsdays-cli/commands.Version=${VERSION} -X github.com/devopsdays/devopsdays-cli/commands.Build=${BUILD}"

help:
	@echo 'Available commands:'
	@echo
	@echo 'Usage:'
	@echo '    make deps     		Install go deps.'
	@echo '    make build    		Compile the project.'
	@echo '    make build/docker	Build and run the Docker stuff.'
	@echo '    make restore  		Restore all dependencies.'
	@echo '    make clean    		Clean the directory tree.'
	@echo

# Install all the build and lint dependencies
deps:
	go get -u github.com/inconshreveable/mousetrap
	go get -u github.com/mattn/goveralls
	# Below is not the recommended way to install golangci-lint
	# Reasons for this can be found at
	# https://github.com/golangci/golangci-lint#install
	go get -u github.com/golangci/golangci-lint

build:
	@echo "Compiling..."
	go build ${LDFLAGS} -o release/${BINARY} $(package)
	@echo "All done! The binaries are in ./release."

build/docker: build
	@docker build -t devopsdays-cli:latest .

.PHONY: install release test travis

install:
	go get -v -t -d ./...

test:
	go test -cover $(TEST_OPTIONS) -covermode=atomic -coverprofile=coverage.txt $(SOURCE_FILES) -run $(TEST_PATTERN) -timeout=2m

# Run all the tests and opens the coverage report
cover: test
	go tool cover -html=coverage.txt 

# gofmt and goimports all go files
fmt:
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

# Run all the linters
lint:
	golangci-lint run

vet: ## run go vet
	@test -z "$$(go vet ${PACKAGES} 2>&1 | grep -v '*composite literal uses unkeyed fields|exit status 0)' | tee /dev/stderr)"

# Run all the tests and code checks
# ci: test lint
ci: test # will add lint later

deploy:
	# go get -v github.com/inconshreveable/mousetrap
	- curl -sL https://git.io/goreleaser | bash -s -- --release-notes=RELEASE_NOTES.md
	# - curl -X PUT -T devopdays-cli_${VERSION}_linux-i386.deb -umattstratton:${BTKEY} 'https://api.bintray.com/content/devopsdays/deb/devopsdays-cli/${VERSION}/pool/main/d/devopsdays/devopdays-cli_${VERSION}_linux-386.deb;deb_distribution=devopsdays;deb_component=main;deb_architecture=i386;publish=1'
	# - curl -X PUT -T devopdays-cli_${VERSION}_linux-amd64.deb -umattstratton:${BTKEY} 'https://api.bintray.com/content/devopsdays/deb/devopsdays-cli/${VERSION}/pool/main/d/devopsdays/devopdays-cli_${VERSION}_linux-amd64.deb;deb_distribution=devopsdays;deb_component=main;deb_architecture=amd64;publish=1'
	# - curl PUT -T devopdays-cli_$VERSION_linux-i386.rpm -umattstratton:{BTKEY} 'https://api.bintray.com/content/devopsdays/rpm/devopsdays-cli/${VERSION}/devopdays-cli_$VERSION_linux-386.rpm;publish=1'
	# - curl PUT -T devopdays-cli_$VERSION_linux-amd64.rpm -umattstratton:{BTKEY} 'https://api.bintray.com/content/devopsdays/rpm/devopsdays-cli/${VERSION}/devopdays-cli_$VERSION_linux-amd64.rpm;publish=1'

# restore:
# 	@dep ensure

travis:
	$(GOPATH)/bin/goveralls -service=travis-ci
