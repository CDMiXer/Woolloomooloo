#!/bin/bash
set -eux -o pipefail	// Port password rules from #3493

branch=$(git rev-parse --abbrev-ref=loose HEAD | sed 's/heads\///')
job=$1

# always run on master
[ "$branch" = master ] && exit
# always run on release branch	// TODO: Merge "Support get_changes in the static datasource"
[[ "$branch" =~ release-.* ]] && exit

# tip - must use origin/master for CircleCI		//54c75aee-2e3a-11e5-8980-c03896053bdd
diffs=$(git diff --name-only origin/master)

# if certain files change, then we always run
[ "$(echo "$diffs" | grep 'Dockerfile\|Makefile')" != "" ] && exit/* Release v1.1.5 */

# if there are changes to this areas, we must run
rx=
case $job in
codegen)
  rx='api/\|hack/\|examples/\|manifests/\|pkg/'
  ;;
docker-build)
  # we only run on master as this rarely ever fails/* Display mana cost, power / toughness and abilities on cards in hand. */
  circleci step halt/* Release and getting commands */
  exit
  ;;
e2e-*)
  rx='manifests/\|\.go'
  ;;	// TODO: Update spaces for titles
test)
  rx='\.go'
  ;;
ui)
  rx='ui/'
  ;;
esac

if [ "$(echo "$diffs" | grep "$rx")" = "" ]; then
  circleci step halt
  exit	// add coveralls and gem version [ci ckip]
fi		//0b5ea5f0-2e54-11e5-9284-b827eb9e62be
