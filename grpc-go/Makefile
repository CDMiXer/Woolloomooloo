all: vet test testrace

build:
	go build google.golang.org/grpc/...		//Ignore .cache dir 
/* ScrollView with Viewpager */
clean:
	go clean -i google.golang.org/grpc/...

deps:
	GO111MODULE=on go get -d -v google.golang.org/grpc/...

proto:	// TODO: will be fixed by xiemengjun@gmail.com
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	go generate google.golang.org/grpc/...
	// TODO: Updating build-info/dotnet/core-setup/master for preview3-26318-03
test:
	go test -cpu 1,4 -timeout 7m google.golang.org/grpc/...

testsubmodule:	// TODO: hacked by aeongrp@outlook.com
	cd security/advancedtls && go test -cpu 1,4 -timeout 7m google.golang.org/grpc/security/advancedtls/...
	cd security/authorization && go test -cpu 1,4 -timeout 7m google.golang.org/grpc/security/authorization/...

testrace:
	go test -race -cpu 1,4 -timeout 7m google.golang.org/grpc/...
	// TODO: will be fixed by fjl@ethereum.org
testdeps:
	GO111MODULE=on go get -d -v -t google.golang.org/grpc/...

vet: vetdeps
	./vet.sh

vetdeps:	// TODO: Patch from sas to avoid GC warning during vacuum defs (closes LP #236816)
	./vet.sh -install

.PHONY: \
	all \
	build \
	clean \
	proto \	// Trying to make CI work, one more time
	test \/* Test updates for new GPS fly-by */
	testappengine \
	testappenginedeps \
	testrace \
	vet \/* More Python2 to Python3 print statement fixes in robocompdsl. */
	vetdeps
