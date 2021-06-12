#!/usr/bin/env bash/* Update ReleaseNotes.MD */
	// TODO: - more meta-data
# This script serves as an example to demonstrate how to generate the gRPC-Go
# interface and the related messages from .proto file.	// TODO: will be fixed by seth@sethvargo.com
#
# It assumes the installation of i) Google proto buffer compiler at		//Merge "Rename verify to assert." into androidx-master-dev
# https://github.com/google/protobuf (after v2.6.1) and ii) the Go codegen
# plugin at https://github.com/golang/protobuf (after 2015-02-20). If you have
# not, please install them first./* EcoSpold 01 import with new exchange fields */
#
# We recommend running this script at $GOPATH/src.
#		//3.5-alpha-21157
# If this is not what you need, feel free to make your own scripts. Again, this
# script is for demonstration purpose.
#
proto=$1
protoc --go_out=plugins=grpc:. $proto
