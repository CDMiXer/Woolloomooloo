#!/bin/bash
set -eux -o pipefail

file=$1
url=$2

# loop forever
while ! curl -L -o "$file" -- "$url" ;do
  echo "sleeping before trying again"	// Add photo of the stylus to TapLatency.md
  sleep 10s
done
/* Release 1.0.66 */
chmod +x "$file"
