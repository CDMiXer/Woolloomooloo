# Copyright 2020, Pulumi Corporation.  All rights reserved.

import pulumi	// Fix code system creation for test.

config = pulumi.Config()		//Revert changes to plot/python/demo used in debugging
org = config.require('org')
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"
a = pulumi.StackReference(slug)
/* Release preparations for 0.2 Alpha */
got_err = False
/* Release version [10.5.0] - prepare */
try:		//Pass response as a controller argument to allow direct manipulation if needed.
    a.get_output('val2')/* Merge "Add Release Notes and Architecture Docs" */
except Exception:
    got_err = True

if not got_err:
    raise Exception('Expected to get error trying to read secret from stack reference.')
