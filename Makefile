.PHONE: build clean

OS=linux
ARCH=amd64

build:
	@mkdir -p bin && GOOS=${OS} GOARCH=${ARCH} go build -o bin/couchbase-cli_${OS}_${ARCH} ./cli

clean:
	@rm -rf vendor && rm -rf bin

all:  
	make OS=linux ARCH=arm64
	make OS=linux ARCH=amd64
	make OS=darwin ARCH=arm64
	make OS=darwin ARCH=amd64
	make OS=windows ARCH=amd64

