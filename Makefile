GO = go

ifndef $(GOPATH)
GOPATH=$(shell go env GOPATH)
export GOPATH
endif

GOCACHE=/tmp
export GOCACHE

BIN_FOLDER=./bin/
BIN_FILE=${BIN_FOLDER}suggest

# Собирает указанный бинарник
build: 
	$(GO) build -o "${BIN_FILE}" ./cmd/
	chmod +x ${BIN_FILE}
	echo "Build finished!"

run: build
	${BIN_FILE}