VERSION?=0.0.0-dev
NAME?=predown
GOOS?=darwin
FLAGS?="-X main.version=${VERSION} -X main.name=${NAME}"
COVERAGE_FILE ?= c.out

build: 
	@ go build -ldflags ${FLAGS} -o dist/${NAME}_${GOOS}

run:
	@ go run -ldflags ${FLAGS} main.go

test:
	@ go test -coverprofile=$(COVERAGE_FILE) $(RACE) ./...

lint:
	@ golint ./..

tool:
	@ go tool cover -$(MODE)=$(COVERAGE_FILE)

race: RACE=-race
race: test

func: MODE=func
func: test tool

html: MODE=html
html: test tool
