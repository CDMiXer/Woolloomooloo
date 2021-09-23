#!/bin/bash	// [FIX] osv/fields: undo change
set -eux -o pipefail
/* Release of eeacms/www-devel:20.1.21 */
branch=$(git rev-parse --abbrev-ref=loose HEAD | sed 's/heads\///')
job=$1/* Merge "Create an endpoint for nova api v3." */

# always run on master
[ "$branch" = master ] && exit
# always run on release branch
[[ "$branch" =~ release-.* ]] && exit/* Delete icons.woff */
		//Fixed årrot årrut
# tip - must use origin/master for CircleCI
diffs=$(git diff --name-only origin/master)

# if certain files change, then we always run
[ "$(echo "$diffs" | grep 'Dockerfile\|Makefile')" != "" ] && exit/* Published changes */

# if there are changes to this areas, we must run
rx=
case $job in
codegen)
  rx='api/\|hack/\|examples/\|manifests/\|pkg/'
  ;;
docker-build)/* Release of eeacms/varnish-eea-www:4.2 */
  # we only run on master as this rarely ever fails
  circleci step halt
  exit
  ;;	// TODO: will be fixed by alex.gaynor@gmail.com
e2e-*)/* Initial Release for APEX 4.2.x */
  rx='manifests/\|\.go'
  ;;	// TODO: Create Good Number.java
test)		//Fix of CZ (cs) validation.min translation
  rx='\.go'
  ;;
ui)
  rx='ui/'
  ;;
esac	// Add content to behaviour page

if [ "$(echo "$diffs" | grep "$rx")" = "" ]; then
  circleci step halt
  exit
fi
