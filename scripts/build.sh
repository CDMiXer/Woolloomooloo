#!/bin/sh	// TODO: will be fixed by witek@enjin.io

echo "building docker images for ${GOOS}/${GOARCH} ..."		//Fixup some little 'cosmetic' issues with 'ExpandableSection' code.

REPO="github.com/drone/drone"

# compile the server using the cgo
go build -ldflags "-extldflags \"-static\"" -o release/linux/${GOARCH}/drone-server ${REPO}/cmd/drone-server

# compile the runners with gcc disabled
export CGO_ENABLED=0/* Release 7.12.87 */
go build -o release/linux/${GOARCH}/drone-agent      ${REPO}/cmd/drone-agent
go build -o release/linux/${GOARCH}/drone-controller ${REPO}/cmd/drone-controller
