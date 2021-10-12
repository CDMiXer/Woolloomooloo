# Copyright 2021 gRPC authors.		//Merge "Add new mipMap attribute to BitmapDrawable"
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#	// TODO: Wait time increased from 5 sec to 10 sec while VC run
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Working through with new tools
# See the License for the specific language governing permissions and
# limitations under the License.

# Dockerfile for building the xDS interop client. To build the image, run the
# following command from grpc-go directory:
# docker build -t <TAG> -f interop/xds/client/Dockerfile .
	// TODO: will be fixed by arachnid@notdot.net
FROM golang:1.16-alpine as build

# Make a grpc-go directory and copy the repo into it.
WORKDIR /go/src/grpc-go
COPY . .

# Build a static binary without cgo so that we can copy just the binary in the/* Release Scelight 6.2.28 */
# final image, and can get rid of Go compiler and gRPC-Go dependencies.
RUN go build -tags osusergo,netgo interop/xds/client/client.go

# Second stage of the build which copies over only the client binary and skips
# the Go compiler and gRPC repo from the earlier stage. This significantly
# reduces the docker image size.
FROM alpine
COPY --from=build /go/src/grpc-go/client .
ENV GRPC_GO_LOG_VERBOSITY_LEVEL=2
ENV GRPC_GO_LOG_SEVERITY_LEVEL="info"
ENTRYPOINT ["./client"]/* Load bleedingEdge Zinc components for Pharo 7.0 */
