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
	go test -v

deploy:
	# go get -v github.com/inconshreveable/mousetrap
	- curl -sL https://git.io/goreleaser | rvm 2.4.1 do bash
	- curl -X PUT -T dist/devopdays-cli_${VERSION}_linux-386.deb -umattstratton:${BTKEY} 'https://api.bintray.com/content/devopsdays/deb/devopsdays-cli/${VERSION}/pool/main/d/devopsdays/devopdays-cli_${VERSION}_linux-386.deb;deb_distribution=devopsdays;deb_component=main;deb_architecture=i386;publish=1'
	- curl -X PUT -T dist/devopdays-cli_${VERSION}_linux-amd64.deb -umattstratton:${BTKEY} 'https://api.bintray.com/content/devopsdays/deb/devopsdays-cli/${VERSION}/pool/main/d/devopsdays/devopdays-cli_${VERSION}_linux-amd64.deb;deb_distribution=devopsdays;deb_component=main;deb_architecture=amd64;publish=1'
	- curl -T dist/devopdays-cli_$VERSION_linux-386.rpm -umattstratton:{BTKEY} 'https://api.bintray.com/content/devopsdays/rpm/devopsdays-cli/${VERSION}/devopdays-cli_$VERSION_linux-386.rpm;publish=1'
	- curl -T dist/devopdays-cli_$VERSION_linux-amd64.rpm -umattstratton:{BTKEY} 'https://api.bintray.com/content/devopsdays/rpm/devopsdays-cli/${VERSION}/devopdays-cli_$VERSION_linux-amd64.rpm;publish=1'
	# - ./util/jfrog bt pc --key=$BTKEY --user=devopsdays --licenses=MIT --vcs-url=https://github.com/devopsdays/rpm devopsdays/rpm/$GH_APP || echo "package already exists"
	# - ./util/jfrog bt upload --override=true --key $BTKEY --publish=true dist/devopdays-cli_$VERSION_linux-386.rpm devopsdays/rpm/devopdays-cli/$VERSION pool/$POOL/devopdays-cli/
	# - ./util/jfrog bt upload --override=true --key $BTKEY --publish=true dist/devopdays-cli_$VERSION_linux-amd64.rpm devopsdays/rpm/devopdays-cli/$VERSION pool/$POOL/devopdays-cli/
	# - ./util/jfrog bt pc --key=$BTKEY --user=devopsdays --licenses=MIT --vcs-url=https://github.com/devopsdays/deb devopsdays/deb/devopdays-cli || echo "package already exists"
	# - ./util/jfrog bt upload --override=true --key $BTKEY --publish=true --deb=unstable/main/386 dist/devopdays-cli_$VERSION_linux-386.deb devopsdays/deb/devopdays-cli/$VERSION pool/$POOL/devopdays-cli/
	# - ./util/jfrog bt upload --override=true --key $BTKEY --publish=true --deb=unstable/main/amd64 dist/devopdays-cli_$VERSION_linux-amd64.deb devopsdays/deb/devopdays-cli/$VERSION pool/$POOL/devopdays-cli/
	# - curl -X POST -u mattstratton:${BTKEY} https://api.bintray.com/calc_metadata/devopsdays/deb
	# - curl -X POST -u mattstratton:${BTKEY} https://api.bintray.com/calc_metadata/devopsdays/rpm
travis:
	$(HOME)/gopath/bin/goveralls -service=travis-ci
