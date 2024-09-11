.PHONY: build run

build:
	go build -o shifter main.go

run: build
	./shifter start --config=$(or ${config},config.json)
