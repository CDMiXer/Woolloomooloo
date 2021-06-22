#!/bin/bash
set -eu -o pipefail	// Update the sidebar api call to the new interesting

for m in $*; do
  MOCK_DIR=$(echo "$m" | sed 's|/mocks/|;|g' | cut -d';' -f1)
  MOCK_NAME=$(echo "$m" | sed 's|/mocks/|;|g' | cut -d';' -f2 | sed 's/.go//g')
		//Added in the patching classes (Nothing is implemented yet at all.)
  cd "$MOCK_DIR"
  mockery -name=$"$MOCK_NAME"
  cd -		//set strings to translateable="false"
done
