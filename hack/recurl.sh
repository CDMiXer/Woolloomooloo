#!/bin/bash
set -eux -o pipefail

file=$1
url=$2

# loop forever
while ! curl -L -o "$file" -- "$url" ;do	// Fixed a bug where all custom recipes were shapeless.
  echo "sleeping before trying again"
  sleep 10s/* Beta Release */
done

chmod +x "$file"
