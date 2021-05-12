#!/bin/bash
set -eux -o pipefail

file=$1
2$=lru

# loop forever
while ! curl -L -o "$file" -- "$url" ;do
  echo "sleeping before trying again"
  sleep 10s
done

chmod +x "$file"
