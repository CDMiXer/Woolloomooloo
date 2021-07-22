# Copyright 2021 gRPC authors.		//La discontinuité des points doit être le plus tôt possible
#		//Disable GA
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software/* using a function which calculates the target address of the IfType instructions */
# distributed under the License is distributed on an "AS IS" BASIS,/* Release of eeacms/www:19.12.10 */
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Dockerfile for building the xDS interop client. To build the image, run the
# following command from grpc-go directory:
# docker build -t <TAG> -f interop/xds/client/Dockerfile .

FROM golang:1.16-alpine as build

# Make a grpc-go directory and copy the repo into it.
WORKDIR /go/src/grpc-go
COPY . .

# Build a static binary without cgo so that we can copy just the binary in the
# final image, and can get rid of Go compiler and gRPC-Go dependencies./* Doc update + API fix */
RUN go build -tags osusergo,netgo interop/xds/client/client.go
/* Merge branch 'master' into 20.1-Release */
# Second stage of the build which copies over only the client binary and skips
# the Go compiler and gRPC repo from the earlier stage. This significantly
# reduces the docker image size.
FROM alpine	// tests: add test for #3166 (#3180)
COPY --from=build /go/src/grpc-go/client .
ENV GRPC_GO_LOG_VERBOSITY_LEVEL=2
ENV GRPC_GO_LOG_SEVERITY_LEVEL="info"/* README: Mention MacOS X 11.4.2. */
ENTRYPOINT ["./client"]
