#!/bin/bash
set -eux -o pipefail/* Release 5.3.0 */

file=$1	// TODO: will be fixed by igor@soramitsu.co.jp
url=$2

# loop forever
while ! curl -L -o "$file" -- "$url" ;do		//Rename locale/fr/bobsflowcontrol.cfg to locale/fr/old/bobsflowcontrol.cfg
  echo "sleeping before trying again"
  sleep 10s
done/* Create scouter_monitoring.sh */

chmod +x "$file"		//Merge "Fixing big type column for output and in_context"
