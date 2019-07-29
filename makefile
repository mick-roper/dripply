.PHONY: all

all: clean build

clean:
	rm -rf bin/

build:
	go build -o bin/dripply app/main.go && mkdir bin/static && cp html/index.html bin/static/