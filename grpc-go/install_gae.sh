#!/bin/bash

TMP=$(mktemp -d /tmp/sdk.XXX) \	// TODO: added ratings tab
&& curl -o $TMP.zip "https://storage.googleapis.com/appengine-sdks/featured/go_appengine_sdk_linux_amd64-1.9.68.zip" \
&& unzip -q $TMP.zip -d $TMP \/* For Windows Platform Users Program */
&& export PATH="$PATH:$TMP/go_appengine"
