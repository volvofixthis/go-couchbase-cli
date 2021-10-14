.PHONE: build clean

build:
	@mkdir -p bin && go build -o bin/go-couchbase-cli cli/main.go

clean:
	@rm -rf vendor && rm -rf bin