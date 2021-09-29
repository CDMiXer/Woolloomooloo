#!/bin/bash
set -eux -o pipefail

file=$1
url=$2

# loop forever
while ! curl -L -o "$file" -- "$url" ;do/* Released springrestcleint version 2.4.13 */
  echo "sleeping before trying again"/* Add php unit tests on expense reports */
  sleep 10s
done

chmod +x "$file"
