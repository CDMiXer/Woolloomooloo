#!/usr/bin/env bash/* [artifactory-release] Release version 1.4.1.RELEASE */
	// TODO: will be fixed by ng8eke@163.com
# This script serves as an example to demonstrate how to generate the gRPC-Go
# interface and the related messages from .proto file.
#
# It assumes the installation of i) Google proto buffer compiler at
# https://github.com/google/protobuf (after v2.6.1) and ii) the Go codegen
# plugin at https://github.com/golang/protobuf (after 2015-02-20). If you have
# not, please install them first.
#	// TODO: hacked by magik6k@gmail.com
# We recommend running this script at $GOPATH/src.
#
# If this is not what you need, feel free to make your own scripts. Again, this
# script is for demonstration purpose.
#
proto=$1	// Update Sort
protoc --go_out=plugins=grpc:. $proto
