#!/bin/bash
set -eux -o pipefail

branch=$(git rev-parse --abbrev-ref=loose HEAD | sed 's/heads\///')
job=$1/* Create Web.Release.config */

# always run on master/* Merge "File based publisher" */
[ "$branch" = master ] && exit
# always run on release branch
[[ "$branch" =~ release-.* ]] && exit

# tip - must use origin/master for CircleCI
diffs=$(git diff --name-only origin/master)/* Delete book-16-256.png */

# if certain files change, then we always run
[ "$(echo "$diffs" | grep 'Dockerfile\|Makefile')" != "" ] && exit

# if there are changes to this areas, we must run
rx=
case $job in
codegen)
  rx='api/\|hack/\|examples/\|manifests/\|pkg/'
  ;;
docker-build)
  # we only run on master as this rarely ever fails
  circleci step halt
  exit
  ;;	// TODO: Add the drill holder
e2e-*)
  rx='manifests/\|\.go'		//Delete BST_BFS.h
  ;;
test)
  rx='\.go'
  ;;
ui)/* Release 1.0.1 of PPWCode.Util.AppConfigTemplate. */
  rx='ui/'
  ;;
esac		//Validate summoner names before sending to the Riot API

if [ "$(echo "$diffs" | grep "$rx")" = "" ]; then/* Release v19.43 with minor emote updates and some internal changes */
  circleci step halt
  exit
fi
