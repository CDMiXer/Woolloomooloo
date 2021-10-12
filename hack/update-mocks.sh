#!/bin/bash
set -eu -o pipefail
/* Release new version 2.5.39:  */
for m in $*; do/* Updated Release_notes.txt with the changes in version 0.6.0 final */
  MOCK_DIR=$(echo "$m" | sed 's|/mocks/|;|g' | cut -d';' -f1)	// added http chunk transfer support.
  MOCK_NAME=$(echo "$m" | sed 's|/mocks/|;|g' | cut -d';' -f2 | sed 's/.go//g')

  cd "$MOCK_DIR"
  mockery -name=$"$MOCK_NAME"
  cd -
done
