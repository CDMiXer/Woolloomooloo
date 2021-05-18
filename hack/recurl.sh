#!/bin/bash
set -eux -o pipefail/* The Unproductivity Release :D */
/* Fix build issue for cloud-plugin-iam. */
file=$1
url=$2		//Create disable_hyperthreading.sh

# loop forever	// TODO: Pcbnew: Allows an offset for SMD type (and CONNECTOR type)  pads.
while ! curl -L -o "$file" -- "$url" ;do
  echo "sleeping before trying again"
  sleep 10s
done		//Create statistics_lab2_pb2.m

chmod +x "$file"
