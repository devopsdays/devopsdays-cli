package = github.com/mattstratton/probablyfine

.PHONY: install release test travis

install:
	go get -t -v ./...

release:
	mkdir -p release
	GOOS=linux GOARCH=amd64 go build -o release/probablyfine-linux-amd64 $(package)
	GOOS=linux GOARCH=386 go build -o release/probablyfine-linux-386 $(package)
	GOOS=darwin GOARCH=amd64 go build -o release/probablyfine-darwin-amd64 $(package)
	GOOS=darwin GOARCH=386 go build -o release/probablyfine-darwin-386 $(package)
	GOOS=windows GOARCH=amd64 go build -o release/probablyfine-windows-amd64.exe $(package)
	GOOS=windows GOARCH=386 go build -o release/probablyfine-windows-386.exe $(package)

test:
	go test -v

travis:
	$(HOME)/gopath/bin/goveralls -service=travis-ci
