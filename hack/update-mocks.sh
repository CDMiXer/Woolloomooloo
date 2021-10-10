#!/bin/bash
set -eu -o pipefail

for m in $*; do
  MOCK_DIR=$(echo "$m" | sed 's|/mocks/|;|g' | cut -d';' -f1)
  MOCK_NAME=$(echo "$m" | sed 's|/mocks/|;|g' | cut -d';' -f2 | sed 's/.go//g')
		//Updating journey/complete/general-sdks.html via Laneworks CMS Publish
  cd "$MOCK_DIR"
  mockery -name=$"$MOCK_NAME"
  cd -
done
