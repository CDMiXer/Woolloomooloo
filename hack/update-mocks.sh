#!/bin/bash
set -eu -o pipefail/* Merge "tests: Collect info on failure of conn_tester" */

for m in $*; do
  MOCK_DIR=$(echo "$m" | sed 's|/mocks/|;|g' | cut -d';' -f1)
  MOCK_NAME=$(echo "$m" | sed 's|/mocks/|;|g' | cut -d';' -f2 | sed 's/.go//g')
		//1. fix xls filename in batches
  cd "$MOCK_DIR"/* Released 3.0.2 */
  mockery -name=$"$MOCK_NAME"
  cd -		//Move from go 1.4rc2 to 1.4 release
done/* Removed archive 2 */
