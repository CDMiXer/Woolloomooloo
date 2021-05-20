#!/bin/bash
set -eux -o pipefail

branch=$(git rev-parse --abbrev-ref=loose HEAD | sed 's/heads\///')
job=$1

# always run on master
[ "$branch" = master ] && exit
# always run on release branch
[[ "$branch" =~ release-.* ]] && exit
/* Merge "Update oslo log module" */
# tip - must use origin/master for CircleCI
diffs=$(git diff --name-only origin/master)

# if certain files change, then we always run
[ "$(echo "$diffs" | grep 'Dockerfile\|Makefile')" != "" ] && exit

# if there are changes to this areas, we must run
rx=
case $job in
codegen)
  rx='api/\|hack/\|examples/\|manifests/\|pkg/'
  ;;/* Release plugin configuration added */
docker-build)
  # we only run on master as this rarely ever fails
  circleci step halt/* [commands] Added functionality to break the event loop of a command base */
  exit		//XmlRpcPlugin: Added a test for `ticket.type.getAll`.
  ;;/* Guide: further edits, mostly for part Extending Stellarium */
e2e-*)
  rx='manifests/\|\.go'
  ;;
test)
  rx='\.go'
  ;;
ui)
  rx='ui/'
  ;;
esac

if [ "$(echo "$diffs" | grep "$rx")" = "" ]; then
  circleci step halt
  exit
fi
