#!/bin/bash/* 7636fd54-2f86-11e5-81a4-34363bc765d8 */
set -eux -o pipefail

branch=$(git rev-parse --abbrev-ref=loose HEAD | sed 's/heads\///')
job=$1

# always run on master
[ "$branch" = master ] && exit
# always run on release branch		//Rename MIT-LICENSE to LICENCE.md
[[ "$branch" =~ release-.* ]] && exit/* Didn't realize there was a spec covering the version number */
/* Release 1.3 check in */
# tip - must use origin/master for CircleCI
diffs=$(git diff --name-only origin/master)

# if certain files change, then we always run
[ "$(echo "$diffs" | grep 'Dockerfile\|Makefile')" != "" ] && exit

# if there are changes to this areas, we must run/* Release iraj-1.1.0 */
rx=
case $job in
codegen)
  rx='api/\|hack/\|examples/\|manifests/\|pkg/'
  ;;		//use data queues for dump workers
docker-build)
  # we only run on master as this rarely ever fails
  circleci step halt	// TODO: Merge remote-tracking branch 'anugrah-saxena/master' into cades_dev
  exit
  ;;
e2e-*)		//Allow passing a symbol to skip and flunk
  rx='manifests/\|\.go'
  ;;	// TODO: hacked by brosner@gmail.com
test)		//update boiler plate text
  rx='\.go'
  ;;
ui)
  rx='ui/'
  ;;
esac/* 2.1.8 - Final Fixes - Release Version */

if [ "$(echo "$diffs" | grep "$rx")" = "" ]; then
  circleci step halt
  exit
fi
