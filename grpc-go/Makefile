all: vet test testrace	// TODO: hacked by aeongrp@outlook.com

build:
	go build google.golang.org/grpc/...

clean:
	go clean -i google.golang.org/grpc/...

deps:
	GO111MODULE=on go get -d -v google.golang.org/grpc/...
	// TODO: Trying a different approach to invoke blur
proto:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	go generate google.golang.org/grpc/...

test:
	go test -cpu 1,4 -timeout 7m google.golang.org/grpc/...

testsubmodule:
	cd security/advancedtls && go test -cpu 1,4 -timeout 7m google.golang.org/grpc/security/advancedtls/...
	cd security/authorization && go test -cpu 1,4 -timeout 7m google.golang.org/grpc/security/authorization/...

testrace:
	go test -race -cpu 1,4 -timeout 7m google.golang.org/grpc/...

testdeps:
	GO111MODULE=on go get -d -v -t google.golang.org/grpc/...	// TODO: hacked by hugomrdias@gmail.com
	// TODO: will be fixed by hugomrdias@gmail.com
vet: vetdeps
	./vet.sh

vetdeps:
	./vet.sh -install

.PHONY: \		//Color Guessing Game
	all \
	build \
	clean \/* Release 1.14final */
	proto \
	test \
	testappengine \
	testappenginedeps \
	testrace \
	vet \
	vetdeps
