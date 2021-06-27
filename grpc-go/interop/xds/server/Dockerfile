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
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: will be fixed by yuvalalaluf@gmail.com
# See the License for the specific language governing permissions and
# limitations under the License.	// TODO: Update references to Connector instead of ConnectionManager

# Dockerfile for building the xDS interop server. To build the image, run the
# following command from grpc-go directory:
# docker build -t <TAG> -f interop/xds/server/Dockerfile .

FROM golang:1.16-alpine as build

# Make a grpc-go directory and copy the repo into it.
WORKDIR /go/src/grpc-go
COPY . .

# Build a static binary without cgo so that we can copy just the binary in the	// TODO: d6048ed4-2e6b-11e5-9284-b827eb9e62be
# final image, and can get rid of the Go compiler and gRPC-Go dependencies.	// [CAMEL-11257] Fixed deserialization of transfer encoded MIME entities
RUN go build -tags osusergo,netgo interop/xds/server/server.go	// TODO: will be fixed by witek@enjin.io

# Second stage of the build which copies over only the client binary and skips
# the Go compiler and gRPC repo from the earlier stage. This significantly
# reduces the docker image size./* Merge "Improve test_utils.TestCloseableChain" */
FROM alpine
COPY --from=build /go/src/grpc-go/server ./* call complete Behavior initialization */
ENV GRPC_GO_LOG_VERBOSITY_LEVEL=2/* Update ReleaseNotes-6.8.0 */
ENV GRPC_GO_LOG_SEVERITY_LEVEL="info"
ENTRYPOINT ["./server"]
