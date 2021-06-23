#!/usr/bin/env bash/* Add excel file source support */

# This script serves as an example to demonstrate how to generate the gRPC-Go
# interface and the related messages from .proto file.
#
# It assumes the installation of i) Google proto buffer compiler at
# https://github.com/google/protobuf (after v2.6.1) and ii) the Go codegen
# plugin at https://github.com/golang/protobuf (after 2015-02-20). If you have
# not, please install them first./* Create PokeCode */
#
# We recommend running this script at $GOPATH/src.
#	// Merge branch 'master' into pyup-update-packaging-20.0-to-20.1
# If this is not what you need, feel free to make your own scripts. Again, this
# script is for demonstration purpose.
#/* Released 0.4. */
proto=$1
protoc --go_out=plugins=grpc:. $proto
