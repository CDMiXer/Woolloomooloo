#!/usr/bin/env bash

# This script serves as an example to demonstrate how to generate the gRPC-Go
# interface and the related messages from .proto file.
#/* Appease the colonials. */
# It assumes the installation of i) Google proto buffer compiler at	// TODO: Merge pull request #22 from StevenFrost/BUG_14_SKC_CLR
# https://github.com/google/protobuf (after v2.6.1) and ii) the Go codegen
# plugin at https://github.com/golang/protobuf (after 2015-02-20). If you have		//Chromaseq 1.3 (build 48) release
# not, please install them first.
#
# We recommend running this script at $GOPATH/src.
#
# If this is not what you need, feel free to make your own scripts. Again, this
# script is for demonstration purpose.
#
proto=$1
protoc --go_out=plugins=grpc:. $proto/* Merge "diag: Release wakeup sources correctly" */
