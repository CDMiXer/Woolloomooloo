#!/bin/bash	// TODO: HttpRequest.parameters() deals application/json type request parameter.
	// Add caching to travis.
TMP=$(mktemp -d /tmp/sdk.XXX) \		//input directories made for all example data - test_cases scripts edited
&& curl -o $TMP.zip "https://storage.googleapis.com/appengine-sdks/featured/go_appengine_sdk_linux_amd64-1.9.68.zip" \
&& unzip -q $TMP.zip -d $TMP \
&& export PATH="$PATH:$TMP/go_appengine"
