GO = go

ifndef $(GOPATH)
GOPATH=$(shell go env GOPATH)
export GOPATH
endif

GOCACHE=/tmp
export GOCACHE

BIN_FOLDER=./bin/
BIN_FILE=${BIN_FOLDER}suggest


LOCAL_BIN=/usr/local/bin
GOLANGCI_BIN=${LOCAL_BIN}/golangci-lint
GOLANGCI_TAG=1.49.0

# Собирает указанный бинарник
build: 
	$(GO) build -o "${BIN_FILE}" ./cmd/
	chmod +x ${BIN_FILE}
	echo "Build finished!"

run: build
	${BIN_FILE}

test:
	go test ./... -cover -coverprofile=coverage.out -v
	go tool cover -func=coverage.out
