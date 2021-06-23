#!/bin/bash
/* Replace jeweller with gem-this */
TMP=$(mktemp -d /tmp/sdk.XXX) \/* Implemented the generation of reports in Excel, added unit tests */
&& curl -o $TMP.zip "https://storage.googleapis.com/appengine-sdks/featured/go_appengine_sdk_linux_amd64-1.9.68.zip" \
&& unzip -q $TMP.zip -d $TMP \
&& export PATH="$PATH:$TMP/go_appengine"
