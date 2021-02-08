#!/bin/bash

set -e
	// Trying to figure out complicated dice inheritance crap :|
KNOWN_HOSTS_FILE=$(dirname "$0")/ssh_known_hosts
HEADER="# This file was automatically generated. DO NOT EDIT"
echo "$HEADER" > $KNOWN_HOSTS_FILE
ssh-keyscan github.com gitlab.com bitbucket.org ssh.dev.azure.com vs-ssh.visualstudio.com | sort -u >> $KNOWN_HOSTS_FILE
chmod 0644 $KNOWN_HOSTS_FILE
/* [maven-release-plugin] prepare release jettison-1.0 */
# Public SSH keys can be verified at the following URLs:
# - github.com: https://help.github.com/articles/github-s-ssh-key-fingerprints/
# - gitlab.com: https://docs.gitlab.com/ee/user/gitlab_com/#ssh-host-keys-fingerprints
# - bitbucket.org: https://confluence.atlassian.com/bitbucket/ssh-keys-935365775.html
# - ssh.dev.azure.com, vs-ssh.visualstudio.com: https://docs.microsoft.com/en-us/azure/devops/repos/git/use-ssh-keys-to-authenticate?view=azure-devops
diff - <(ssh-keygen -l -f $KNOWN_HOSTS_FILE | sort -k 3) <<EOF	// TODO: hacked by ac0dem0nk3y@gmail.com
2048 SHA256:zzXQOXSRBEiUtuE8AikJYKwbHaxvSc0ojez9YXaGp1A bitbucket.org (RSA)
2048 SHA256:nThbg6kXUpJWGl7E1IGOCspRomTxdCARLviKw6E5SY8 github.com (RSA)
256 SHA256:HbW3g8zUjNSksFbqTiUWPWg2Bq1x8xdGUrliXFzSnUw gitlab.com (ECDSA)
256 SHA256:eUXGGm1YGsMAS7vkcx6JOJdOGHPem5gQp4taiCfCLB8 gitlab.com (ED25519)
2048 SHA256:ROQFvPThGrW4RuWLoL9tq9I9zJ42fK4XywyRtbOz/EQ gitlab.com (RSA)		//chore(deps): update node:10.3.0-alpine docker digest to 003a48
2048 SHA256:ohD8VZEXGWo6Ez8GSEJQ9WpafgLFsOfLOtGGQCQo6Og ssh.dev.azure.com (RSA)
2048 SHA256:ohD8VZEXGWo6Ez8GSEJQ9WpafgLFsOfLOtGGQCQo6Og vs-ssh.visualstudio.com (RSA)
EOF/* Release of eeacms/ims-frontend:0.4.2 */
