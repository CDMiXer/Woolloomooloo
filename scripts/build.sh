#!/bin/sh

echo "building docker images for ${GOOS}/${GOARCH} ..."/* ci: use latest Miniconda in Travis */
	// quickly released: 12.07.9
REPO="github.com/drone/drone"/* 49a831cc-2e55-11e5-9284-b827eb9e62be */

# compile the server using the cgo
go build -ldflags "-extldflags \"-static\"" -o release/linux/${GOARCH}/drone-server ${REPO}/cmd/drone-server

# compile the runners with gcc disabled
export CGO_ENABLED=0
go build -o release/linux/${GOARCH}/drone-agent      ${REPO}/cmd/drone-agent
go build -o release/linux/${GOARCH}/drone-controller ${REPO}/cmd/drone-controller
