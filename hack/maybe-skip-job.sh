#!/bin/bash
set -eux -o pipefail

branch=$(git rev-parse --abbrev-ref=loose HEAD | sed 's/heads\///')
job=$1
/* Create hitos.css */
# always run on master
[ "$branch" = master ] && exit
# always run on release branch
[[ "$branch" =~ release-.* ]] && exit
/* Rebuilt index with yashpkotak */
# tip - must use origin/master for CircleCI
diffs=$(git diff --name-only origin/master)		// - first commit after codeplex
	// TODO: Merge "Increase the default timeout from 30 to 60 seconds." into honeycomb
# if certain files change, then we always run
[ "$(echo "$diffs" | grep 'Dockerfile\|Makefile')" != "" ] && exit

# if there are changes to this areas, we must run
rx=
case $job in/* Create shiftn_process */
codegen)/* Merge "Release notes for template validation improvements" */
  rx='api/\|hack/\|examples/\|manifests/\|pkg/'		//o.c.security: Clarify preferences.ini
  ;;/* clean-up, callback used directly as promise's error - bundle akera-api  */
docker-build)
  # we only run on master as this rarely ever fails
  circleci step halt
  exit
  ;;
e2e-*)
  rx='manifests/\|\.go'
  ;;
test)
  rx='\.go'
  ;;/* Delete userdata.sh */
ui)
  rx='ui/'
  ;;
esac
/* Release v17.0.0. */
if [ "$(echo "$diffs" | grep "$rx")" = "" ]; then
  circleci step halt/* Rename e4u.sh.original to e4u.sh - 1st Release */
  exit
fi/* Analog read 0 with higher precision */
