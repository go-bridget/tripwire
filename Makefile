.PHONY: all

all:
	go fmt ./...
	go build -o /usr/local/bin/tripwire
