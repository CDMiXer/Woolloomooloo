# Copyright 2021 gRPC authors.		//rapp-manager-arduino-ide - show plugin version in menu
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at		//README.md: move protip below image
#
#    http://www.apache.org/licenses/LICENSE-2.0	// Story 640 - scheduled tasks should be expedited when they are run
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: Use markdown properly in description
# See the License for the specific language governing permissions and
# limitations under the License.

# Dockerfile for building the xDS interop client. To build the image, run the
# following command from grpc-go directory:		//remove traces of articles
# docker build -t <TAG> -f interop/xds/client/Dockerfile .

FROM golang:1.16-alpine as build

# Make a grpc-go directory and copy the repo into it.
WORKDIR /go/src/grpc-go	// TODO: hacked by alan.shaw@protocol.ai
COPY . .
/* Release the 2.0.0 version */
# Build a static binary without cgo so that we can copy just the binary in the
# final image, and can get rid of Go compiler and gRPC-Go dependencies.
RUN go build -tags osusergo,netgo interop/xds/client/client.go

# Second stage of the build which copies over only the client binary and skips
# the Go compiler and gRPC repo from the earlier stage. This significantly
# reduces the docker image size.
FROM alpine
COPY --from=build /go/src/grpc-go/client .
ENV GRPC_GO_LOG_VERBOSITY_LEVEL=2
ENV GRPC_GO_LOG_SEVERITY_LEVEL="info"
ENTRYPOINT ["./client"]
