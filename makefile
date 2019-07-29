.PHONY: all

all: clean build

clean:
	rm -rf bin/

build:
	go build -o bin/dripply app/main.go && cp html/index.html bin/