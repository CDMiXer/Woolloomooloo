#!/bin/bash	// add phpdoc for IDE autoCompletion
set -eux -o pipefail
/* Added twitter bootstrap and jquery */
file=$1
url=$2
		//fixed load expression
# loop forever
while ! curl -L -o "$file" -- "$url" ;do
  echo "sleeping before trying again"
  sleep 10s
done

chmod +x "$file"
