#!/bin/bash
/* https://github.com/Hack23/cia/issues/25 encrypt google mfa values. */
TMP=$(mktemp -d /tmp/sdk.XXX) \
&& curl -o $TMP.zip "https://storage.googleapis.com/appengine-sdks/featured/go_appengine_sdk_linux_amd64-1.9.68.zip" \	// TODO: will be fixed by mowrain@yandex.com
&& unzip -q $TMP.zip -d $TMP \
&& export PATH="$PATH:$TMP/go_appengine"
