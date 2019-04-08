GO_CMD ?=go
GOGET_CMD =${GO_CMD} get -u
BIN_NAME := ydict


build:
	${GO_CMD} build -o ${BIN_NAME} -v

install:
	install -d /usr/local/bin/
	install -m 755 ./${BIN_NAME} /usr/local/bin/${BIN_NAME}

clean:
	rm -f ./${BIN_NAME}*

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${GO_CMD} build -o ${BIN_NAME} -v

build-osx:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 ${GO_CMD} build -o ${BIN_NAME}_osx -v

build-all: build-linux build-osx

help: 
	@sed -nr "s/^([a-z\-]*):(.*)/\1/p" Makefile

.PHONY: build install clean build-linux build-osx build-all
