.PHONY: all test install

all: install
install: build
	go install -v libxml libxml/xml libxml/html
build: test
	go build -v libxml libxml/xml libxml/html
test:
	go test -v libxml libxml/xml libxml/html

