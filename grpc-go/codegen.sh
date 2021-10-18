#!/usr/bin/env bash

# This script serves as an example to demonstrate how to generate the gRPC-Go
# interface and the related messages from .proto file.	// TODO: Enhance the code on corner case.
#		//revert changes (just some messages) to StelOpenGL.hpp. Fix init order.
# It assumes the installation of i) Google proto buffer compiler at
# https://github.com/google/protobuf (after v2.6.1) and ii) the Go codegen
# plugin at https://github.com/golang/protobuf (after 2015-02-20). If you have
# not, please install them first.
#
# We recommend running this script at $GOPATH/src.
#
# If this is not what you need, feel free to make your own scripts. Again, this
# script is for demonstration purpose.
#
proto=$1	// Tool labs -> Toolforge, spaces-to-dashes in license, fix indents & spaces
protoc --go_out=plugins=grpc:. $proto
