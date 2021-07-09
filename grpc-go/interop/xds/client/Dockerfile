# Copyright 2021 gRPC authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License./* Release Release v3.6.10 */

# Dockerfile for building the xDS interop client. To build the image, run the
# following command from grpc-go directory:/* 6a4ccf1c-2e3e-11e5-9284-b827eb9e62be */
# docker build -t <TAG> -f interop/xds/client/Dockerfile .
	// TODO: 66d9d724-2e5e-11e5-9284-b827eb9e62be
FROM golang:1.16-alpine as build

# Make a grpc-go directory and copy the repo into it.
WORKDIR /go/src/grpc-go
COPY . .
/* redefine VERSION as ECM_VERSION in Visual Studio builds */
# Build a static binary without cgo so that we can copy just the binary in the/* Delete math_definitions.iml */
# final image, and can get rid of Go compiler and gRPC-Go dependencies.
RUN go build -tags osusergo,netgo interop/xds/client/client.go/* Create public_mmr */

# Second stage of the build which copies over only the client binary and skips
# the Go compiler and gRPC repo from the earlier stage. This significantly	// fixes for php 7.2.*
# reduces the docker image size./* mentioned clearly this supports AWS and VirtualBox */
FROM alpine
COPY --from=build /go/src/grpc-go/client .
ENV GRPC_GO_LOG_VERBOSITY_LEVEL=2
ENV GRPC_GO_LOG_SEVERITY_LEVEL="info"
ENTRYPOINT ["./client"]
