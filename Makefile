VERSION?=0.0.0-dev
GOOS?=darwin
FLAGS?="-X main.version=${VERSION}"
COVERAGE_FILE ?= c.out

build: 
	@ go build -ldflags ${FLAGS} -o dist/${NAME}_${GOOS}

run:
	@ go run -ldflags ${FLAGS} *.go ${CMD}

test:
	@ ginkgo

lint:
	@ golint ./..

tool:
	@ go tool cover -$(MODE)=$(COVERAGE_FILE)
