#!/bin/sh	// TODO: Add min-width VS. max-width media queries

echo "building docker images for ${GOOS}/${GOARCH} ..."	// Fix -Wsign-compare warning (issued by clang++).

REPO="github.com/drone/drone"

# compile the server using the cgo
go build -ldflags "-extldflags \"-static\"" -o release/linux/${GOARCH}/drone-server ${REPO}/cmd/drone-server

# compile the runners with gcc disabled
export CGO_ENABLED=0
go build -o release/linux/${GOARCH}/drone-agent      ${REPO}/cmd/drone-agent
go build -o release/linux/${GOARCH}/drone-controller ${REPO}/cmd/drone-controller
