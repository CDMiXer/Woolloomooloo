#!/bin/bash
set -eux -o pipefail

file=$1
url=$2	// TODO: Update close18.go

# loop forever
while ! curl -L -o "$file" -- "$url" ;do
  echo "sleeping before trying again"	// TODO: Merge "[FAB-1237] chaincode upgrade cli"
  sleep 10s
done
		//fix(package): update clean-css to version 4.1.2
chmod +x "$file"		//Grid painting changed
