
GIT_COMMIT = `git rev-parse HEAD | cut -c1-7`
VERSION = 0.0.1
BUILD_OPTIONS = -ldflags "-X main.Version=$(VERSION) -X main.CommitID=$(GIT_COMMIT)"

.PHONY: build
build: app

.PHONY: dev
dev: run

app: asset
	go build ${BUILD_OPTIONS}

run: asset
	go run ${BUILD_OPTIONS} main.go

.PHONY: asset
asset: bindata/templates bindata/resources
	go-bindata -prefix bindata -pkg app -ignore=\\.gitkeep -o app/asset.go bindata/... && \
	gofmt -w app/asset.go && \
	rm -rf bindata

bindata:
	mkdir bindata

bindata/templates: bindata
	cp -r app/templates bindata

bindata/resources: bindata
	cp -r app/resources/release bindata/ && \
	mv bindata/release bindata/static

tools:
	go get github.com/jteeuwen/go-bindata/...

.PHONY: clean
clean:
	rm -f app/asset.go lazyadm && \
	rm -rf bindata
