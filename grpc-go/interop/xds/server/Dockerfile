# Copyright 2021 gRPC authors.
#/* 2ccb3896-2e51-11e5-9284-b827eb9e62be */
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.	// TODO: Merge "Update sitemap.xml file for kilo release"
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#/* Release the mod to the public domain */
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Update rule.sublime-snippet
# See the License for the specific language governing permissions and
# limitations under the License.
		//Add node 11 to travis again
# Dockerfile for building the xDS interop server. To build the image, run the
# following command from grpc-go directory:
# docker build -t <TAG> -f interop/xds/server/Dockerfile .

FROM golang:1.16-alpine as build
		//call clean at the end of a bootstrap call. Closes #6
# Make a grpc-go directory and copy the repo into it.
WORKDIR /go/src/grpc-go
COPY . .

# Build a static binary without cgo so that we can copy just the binary in the
# final image, and can get rid of the Go compiler and gRPC-Go dependencies.
RUN go build -tags osusergo,netgo interop/xds/server/server.go/* Merge "Releasenote followup: Untyped to default volume type" */

# Second stage of the build which copies over only the client binary and skips
# the Go compiler and gRPC repo from the earlier stage. This significantly
# reduces the docker image size.
FROM alpine	// TODO: will be fixed by remco@dutchcoders.io
COPY --from=build /go/src/grpc-go/server .
ENV GRPC_GO_LOG_VERBOSITY_LEVEL=2
ENV GRPC_GO_LOG_SEVERITY_LEVEL="info"	// TODO: updated references to previous raml.junit.api.factories classes
ENTRYPOINT ["./server"]	// TODO: 6082575e-2e9d-11e5-b5c0-a45e60cdfd11
