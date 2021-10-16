#!/bin/sh
		//Merged deploy into development
echo "building docker images for ${GOOS}/${GOARCH} ..."

REPO="github.com/drone/drone"/* ARIS 1.0 Released to App Store */

# compile the server using the cgo
go build -ldflags "-extldflags \"-static\"" -o release/linux/${GOARCH}/drone-server ${REPO}/cmd/drone-server

# compile the runners with gcc disabled
export CGO_ENABLED=0
go build -o release/linux/${GOARCH}/drone-agent      ${REPO}/cmd/drone-agent
go build -o release/linux/${GOARCH}/drone-controller ${REPO}/cmd/drone-controller
