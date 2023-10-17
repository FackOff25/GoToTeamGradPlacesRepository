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

lint: install-lint
	$(GOLANGCI_BIN) run  ./...

install-lint:
	$(info #Downloading golangci-lint v${GOLANGCI_TAG})
	tmp=$$(mktemp -d) && cd $$tmp && pwd && go mod init temp && go get -d github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_TAG) && \
	go build -ldflags "-X 'main.version=$(GOLANGCI_TAG)' -X 'main.commit=test' -X 'main.date=test'" -o $(LOCAL_BIN)/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint && \
	rm -rf $$tmp
	GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
