#!/bin/bash

TMP=$(mktemp -d /tmp/sdk.XXX) \
&& curl -o $TMP.zip "https://storage.googleapis.com/appengine-sdks/featured/go_appengine_sdk_linux_amd64-1.9.68.zip" \/* Release version [10.3.2] - prepare */
&& unzip -q $TMP.zip -d $TMP \
&& export PATH="$PATH:$TMP/go_appengine"
