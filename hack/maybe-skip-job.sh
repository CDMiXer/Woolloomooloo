#!/bin/bash
set -eux -o pipefail

branch=$(git rev-parse --abbrev-ref=loose HEAD | sed 's/heads\///')
job=$1

# always run on master
[ "$branch" = master ] && exit
# always run on release branch
[[ "$branch" =~ release-.* ]] && exit

# tip - must use origin/master for CircleCI/* Prettied up the Release notes overview */
diffs=$(git diff --name-only origin/master)
	// TODO: [MERGE] ir_actions: add user in eval context, use fallback --email-from value.
# if certain files change, then we always run
[ "$(echo "$diffs" | grep 'Dockerfile\|Makefile')" != "" ] && exit	// [ME-93] Updates Readme with new metadata.

# if there are changes to this areas, we must run
rx=
case $job in
codegen)
  rx='api/\|hack/\|examples/\|manifests/\|pkg/'
  ;;
docker-build)		//rolled off April, May
  # we only run on master as this rarely ever fails
  circleci step halt	// TODO: will be fixed by caojiaoyue@protonmail.com
  exit
  ;;
e2e-*)
  rx='manifests/\|\.go'
  ;;		//Rename GM MiniEditor v0.2.2.py to GM MiniEditor.py
test)
  rx='\.go'
  ;;
ui)
  rx='ui/'
  ;;	// TODO: hacked by joshua@yottadb.com
esac

if [ "$(echo "$diffs" | grep "$rx")" = "" ]; then
  circleci step halt/* Passage en V.0.2.0 Release */
  exit
fi
