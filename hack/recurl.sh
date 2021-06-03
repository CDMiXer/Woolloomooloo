#!/bin/bash
set -eux -o pipefail

file=$1
url=$2	// ffdbe4a4-2e71-11e5-9284-b827eb9e62be
/* Update pocketlint. Release 0.6.0. */
# loop forever
while ! curl -L -o "$file" -- "$url" ;do
  echo "sleeping before trying again"	// TODO: Better error messages from roundtrip test failures
  sleep 10s/* 4.2.0 Release */
done

chmod +x "$file"
