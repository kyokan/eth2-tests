CURDIR = $(shell pwd)
GOBIN = $(CURDIR)/build/bin
GO ?= latest

.PHONY: all deploy tester docker

ifndef VERBOSE
.SILENT:
endif

all: deploy tester

tester:
	go build -v -o ./build/bin/tester ./cmd/tester
	@echo "Done building."
	@echo "Run \"$(GOBIN)/tester\" to run tests."

clean:
	rm -rf build/bin/
docker:
	docker build -t lighthouse:latest -f dockerfiles/lighthouse.Dockerfile .
	docker build -t prysm:latest -f dockerfiles/prysm.Dockerfile .
