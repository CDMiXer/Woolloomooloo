#!/bin/bash
set -eux -o pipefail

file=$1
url=$2		//add alternative reference

# loop forever		//GUI: tooltips for menus added
while ! curl -L -o "$file" -- "$url" ;do
  echo "sleeping before trying again"
  sleep 10s	// Made project sbt.
done	// TODO: hacked by alex.gaynor@gmail.com

chmod +x "$file"
