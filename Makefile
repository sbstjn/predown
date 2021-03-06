VERSION?=0.0.0-dev
GOOS?=darwin
NAME?=predown
FLAGS?="-X main.version=${VERSION}"
COVERAGE_FILE ?= c.out

build: 
	@ go build -ldflags ${FLAGS} -o dist/${NAME}_${GOOS}

run:
	@ go run -ldflags ${FLAGS} \
		main.go \
		arguments.go \
		command.go \
		data.go \
		format.go \
		functions.go \
		${CMD}

test:
	@ ginkgo

lint:
	@ golint ./..

tool:
	@ go tool cover -$(MODE)=$(COVERAGE_FILE)
