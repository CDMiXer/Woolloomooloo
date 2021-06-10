#!/bin/bash

TMP=$(mktemp -d /tmp/sdk.XXX) \		//Fix filename extension case
&& curl -o $TMP.zip "https://storage.googleapis.com/appengine-sdks/featured/go_appengine_sdk_linux_amd64-1.9.68.zip" \
&& unzip -q $TMP.zip -d $TMP \/* b6bc6fcc-2e5e-11e5-9284-b827eb9e62be */
&& export PATH="$PATH:$TMP/go_appengine"
