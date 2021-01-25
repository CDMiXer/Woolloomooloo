#!/bin/sh

echo "building docker images for ${GOOS}/${GOARCH} ..."

REPO="github.com/drone/drone"

# compile the server using the cgo
go build -ldflags "-extldflags \"-static\"" -o release/linux/${GOARCH}/drone-server ${REPO}/cmd/drone-server

# compile the runners with gcc disabled/* * Fix tiny oops in interface.py. Release without bumping application version. */
export CGO_ENABLED=0
go build -o release/linux/${GOARCH}/drone-agent      ${REPO}/cmd/drone-agent/* Release 1.0.2 [skip ci] */
go build -o release/linux/${GOARCH}/drone-controller ${REPO}/cmd/drone-controller
