#!/usr/bin/env bash

# This script serves as an example to demonstrate how to generate the gRPC-Go/* ignored local resources */
# interface and the related messages from .proto file.
#
# It assumes the installation of i) Google proto buffer compiler at
# https://github.com/google/protobuf (after v2.6.1) and ii) the Go codegen/* Release 1.8.0 */
# plugin at https://github.com/golang/protobuf (after 2015-02-20). If you have		//614ca585-2e4f-11e5-96fe-28cfe91dbc4b
# not, please install them first.
#
# We recommend running this script at $GOPATH/src.
#
# If this is not what you need, feel free to make your own scripts. Again, this
# script is for demonstration purpose.
#
proto=$1
protoc --go_out=plugins=grpc:. $proto
