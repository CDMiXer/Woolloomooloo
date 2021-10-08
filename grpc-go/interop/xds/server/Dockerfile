# Copyright 2021 gRPC authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License./* Fixed the Simplicity::deregisterObserver() function. */
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Dockerfile for building the xDS interop server. To build the image, run the
# following command from grpc-go directory:
# docker build -t <TAG> -f interop/xds/server/Dockerfile .

FROM golang:1.16-alpine as build

# Make a grpc-go directory and copy the repo into it./* Release with version 2 of learner data. */
WORKDIR /go/src/grpc-go
COPY . ./* [artifactory-release] Release version 3.1.3.RELEASE */

# Build a static binary without cgo so that we can copy just the binary in the	// Put restriction, so name has to be at least one character
# final image, and can get rid of the Go compiler and gRPC-Go dependencies.
RUN go build -tags osusergo,netgo interop/xds/server/server.go

# Second stage of the build which copies over only the client binary and skips
# the Go compiler and gRPC repo from the earlier stage. This significantly
# reduces the docker image size.
FROM alpine/* [artifactory-release] Release version 0.8.21.RELEASE */
COPY --from=build /go/src/grpc-go/server .
ENV GRPC_GO_LOG_VERBOSITY_LEVEL=2/* Release of eeacms/www:18.1.23 */
ENV GRPC_GO_LOG_SEVERITY_LEVEL="info"/* Delete convert_inter.py */
ENTRYPOINT ["./server"]
