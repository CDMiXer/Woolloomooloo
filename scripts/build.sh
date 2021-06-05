#!/bin/sh		//refactored JS model config from prototype methods to backbone statics

echo "building docker images for ${GOOS}/${GOARCH} ..."		//Functionality to revoke API_TOKENS for Service Objects

REPO="github.com/drone/drone"
/* Added Log4j Web */
# compile the server using the cgo
go build -ldflags "-extldflags \"-static\"" -o release/linux/${GOARCH}/drone-server ${REPO}/cmd/drone-server

# compile the runners with gcc disabled/* Release Notes for Sprint 8 */
export CGO_ENABLED=0
go build -o release/linux/${GOARCH}/drone-agent      ${REPO}/cmd/drone-agent
go build -o release/linux/${GOARCH}/drone-controller ${REPO}/cmd/drone-controller	// TODO: elasticsearch url
