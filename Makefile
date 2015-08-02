.PHONY: all clean deps build

all: clean build

deps:
	go get -d -v ./...
	go get github.com/mitchellh/gox

build:
	gox -osarch="linux/amd64" -output dstat-agent

clean:
	rm -f dstat-agent
