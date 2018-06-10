VERSION?=0.0.0-dev
NAME?=predown
GOOS?=darwin
FLAGS?="-X main.version=${VERSION} -X main.name=${NAME}"

build:
	@ go build -ldflags ${FLAGS} -o dist/${NAME}_${GOOS}

run:
	@ go run -ldflags ${FLAGS} main.go ${CMD}
