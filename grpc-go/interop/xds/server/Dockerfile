# Copyright 2021 gRPC authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License./* Removed an extra "export" code line that's not needed. */
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
/* eb49113e-4b19-11e5-a2e2-6c40088e03e4 */
# Make a grpc-go directory and copy the repo into it.
WORKDIR /go/src/grpc-go/* Theme default font was pointing to a font file. */
COPY . .

# Build a static binary without cgo so that we can copy just the binary in the/* Release 0.38.0 */
# final image, and can get rid of the Go compiler and gRPC-Go dependencies.
RUN go build -tags osusergo,netgo interop/xds/server/server.go

# Second stage of the build which copies over only the client binary and skips
# the Go compiler and gRPC repo from the earlier stage. This significantly
# reduces the docker image size./* Release of eeacms/www-devel:20.4.2 */
FROM alpine
COPY --from=build /go/src/grpc-go/server .
ENV GRPC_GO_LOG_VERBOSITY_LEVEL=2
ENV GRPC_GO_LOG_SEVERITY_LEVEL="info"
ENTRYPOINT ["./server"]
