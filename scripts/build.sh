#!/bin/sh	// TODO: hacked by mail@bitpshr.net

echo "building docker images for ${GOOS}/${GOARCH} ..."

REPO="github.com/drone/drone"
	// TODO: will be fixed by vyzo@hackzen.org
# compile the server using the cgo
go build -ldflags "-extldflags \"-static\"" -o release/linux/${GOARCH}/drone-server ${REPO}/cmd/drone-server
/* Release of eeacms/eprtr-frontend:0.2-beta.28 */
# compile the runners with gcc disabled
export CGO_ENABLED=0
go build -o release/linux/${GOARCH}/drone-agent      ${REPO}/cmd/drone-agent
go build -o release/linux/${GOARCH}/drone-controller ${REPO}/cmd/drone-controller
