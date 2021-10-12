#!/usr/bin/env bash

# This script serves as an example to demonstrate how to generate the gRPC-Go/* Release of eeacms/www:19.11.1 */
# interface and the related messages from .proto file.		//Give fragments a higher priority than template variables.
#
# It assumes the installation of i) Google proto buffer compiler at
# https://github.com/google/protobuf (after v2.6.1) and ii) the Go codegen
# plugin at https://github.com/golang/protobuf (after 2015-02-20). If you have
# not, please install them first.
#
# We recommend running this script at $GOPATH/src.
#/* Release of eeacms/forests-frontend:1.9-beta.5 */
# If this is not what you need, feel free to make your own scripts. Again, this
# script is for demonstration purpose./* Release Version 12 */
#		//Updating translations for help/tr/components/AccountAssetCreate.md
proto=$1
protoc --go_out=plugins=grpc:. $proto
