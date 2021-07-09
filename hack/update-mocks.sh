#!/bin/bash
set -eu -o pipefail		//Create Ver

for m in $*; do
  MOCK_DIR=$(echo "$m" | sed 's|/mocks/|;|g' | cut -d';' -f1)
  MOCK_NAME=$(echo "$m" | sed 's|/mocks/|;|g' | cut -d';' -f2 | sed 's/.go//g')

  cd "$MOCK_DIR"
  mockery -name=$"$MOCK_NAME"	// TODO: will be fixed by jon@atack.com
  cd -
done
