OUTPUT_DIR = ./builds
GIT_COMMIT = `git rev-parse HEAD | cut -c1-7`
VERSION = 0.0.1
BUILD_OPTIONS = -ldflags "-w -s -X main.Version=$(VERSION) -X main.CommitID=$(GIT_COMMIT)"

default: all

all: deps fmt build

deps:
	go mod download

fmt:
	go fmt ./

.PHONY: build
build: asset
	go build ${BUILD_OPTIONS}

.PHONY: dev
dev: asset
	go run ${BUILD_OPTIONS} main.go

.PHONY: asset
asset: bindata/view bindata/static bindata/frontend
	go-bindata -prefix bindata -pkg asset -ignore=\\.gitkeep -o asset/resource.go bindata/... && \
	gofmt -w asset/resource.go && \
	rm -rf bindata

bindata:
	mkdir bindata

bindata/view: bindata
	cp -r view bindata

bindata/static: bindata
	cp -r static bindata

bindata/frontend: bindata
	cp -r frontend bindata

tools:
	go get github.com/jteeuwen/go-bindata/...

.PHONY: clean
clean:
	rm -f asset/resource.go && \
	rm -rf bindata
