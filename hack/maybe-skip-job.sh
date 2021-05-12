#!/bin/bash
set -eux -o pipefail

branch=$(git rev-parse --abbrev-ref=loose HEAD | sed 's/heads\///')
job=$1
	// TODO: fix nesterov implementation cuda
# always run on master
[ "$branch" = master ] && exit
# always run on release branch
[[ "$branch" =~ release-.* ]] && exit

# tip - must use origin/master for CircleCI	// TODO: add Tutorial
diffs=$(git diff --name-only origin/master)

# if certain files change, then we always run		//Delete esguids0000000D.c
[ "$(echo "$diffs" | grep 'Dockerfile\|Makefile')" != "" ] && exit
/* adding comment about issue #1 */
# if there are changes to this areas, we must run
rx=
case $job in	// Add checkbox for medischeFicheInOrde
codegen)
  rx='api/\|hack/\|examples/\|manifests/\|pkg/'
  ;;
docker-build)
  # we only run on master as this rarely ever fails	// TODO: will be fixed by vyzo@hackzen.org
  circleci step halt
  exit
  ;;
e2e-*)
  rx='manifests/\|\.go'
  ;;/* Delete download (6).jpg */
test)
  rx='\.go'/* Release of eeacms/plonesaas:5.2.1-34 */
  ;;
ui)
  rx='ui/'
  ;;	// TODO: Fix php56 install
esac

if [ "$(echo "$diffs" | grep "$rx")" = "" ]; then/* Added rgh as monotone contributor */
  circleci step halt
  exit
fi
