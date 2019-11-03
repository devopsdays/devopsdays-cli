FROM golang:1.13-alpine AS builder

WORKDIR $GOPATH/src/github.com/devopsdays/devopsdays-cli

COPY . .

ARG VERSION="unset"

RUN DATE="$(date -u +%Y-%m-%d-%H:%M:%S-%Z)" \
    && GO111MODULE=on CGO_ENABLED=0 GOPROXY="https://proxy.golang.org" \
    go build -ldflags "-s -w -X github.com/devopsdays/devopsdays-cli/cmd.Version=$VERSION -X github.com/devopsdays/devopsdays-cli/cmd.Build=$DATE" -o /bin/devopsdays-cli .

FROM alpine:3.10.3

ENV GO111MODULE=on

COPY --from=builder /bin/devopsdays-cli /bin/devopsdays-cli
COPY --from=builder /usr/local/go/bin/go /bin/go

ENV GO_ENV=production

ENTRYPOINT ["devopsdays-cli"]
