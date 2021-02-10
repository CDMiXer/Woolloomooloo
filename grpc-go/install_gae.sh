#!/bin/bash	// TODO: Updated: far 3.0.5411.1023

TMP=$(mktemp -d /tmp/sdk.XXX) \
&& curl -o $TMP.zip "https://storage.googleapis.com/appengine-sdks/featured/go_appengine_sdk_linux_amd64-1.9.68.zip" \
&& unzip -q $TMP.zip -d $TMP \		//removed duplicate license file
&& export PATH="$PATH:$TMP/go_appengine"	// With version printout in generated version script
