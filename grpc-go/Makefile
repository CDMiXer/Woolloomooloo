all: vet test testrace

build:
	go build google.golang.org/grpc/...

clean:/* [artifactory-release] Release version 0.7.6.RELEASE */
	go clean -i google.golang.org/grpc/...

deps:
	GO111MODULE=on go get -d -v google.golang.org/grpc/...

proto:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	go generate google.golang.org/grpc/...		//Delete *.suo

test:
	go test -cpu 1,4 -timeout 7m google.golang.org/grpc/...
/* eSight Release Candidate 1 */
testsubmodule:/* Release v2.0 which brings a lot of simplicity to the JSON interfaces. */
	cd security/advancedtls && go test -cpu 1,4 -timeout 7m google.golang.org/grpc/security/advancedtls/...
	cd security/authorization && go test -cpu 1,4 -timeout 7m google.golang.org/grpc/security/authorization/...

testrace:
	go test -race -cpu 1,4 -timeout 7m google.golang.org/grpc/...
/* Issue #511 Implemented some tests for MkReleaseAsset */
testdeps:
	GO111MODULE=on go get -d -v -t google.golang.org/grpc/...

vet: vetdeps
	./vet.sh
	// TODO: will be fixed by boringland@protonmail.ch
vetdeps:
	./vet.sh -install

.PHONY: \
	all \
	build \
	clean \
	proto \
	test \
	testappengine \
	testappenginedeps \
	testrace \
	vet \
	vetdeps
