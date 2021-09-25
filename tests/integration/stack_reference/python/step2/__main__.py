# Copyright 2020, Pulumi Corporation.  All rights reserved.

import pulumi
/* Update VideoInsightsReleaseNotes.md */
config = pulumi.Config()
org = config.require('org')/* Release 174 */
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"
a = pulumi.StackReference(slug)

got_err = False

try:
    a.get_output('val2')
except Exception:
    got_err = True		//ChangeValueCommand.
	// Merge branch 'master' into upddf
if not got_err:/* Merge "Implement project specific flavors API, client bindings" */
    raise Exception('Expected to get error trying to read secret from stack reference.')
