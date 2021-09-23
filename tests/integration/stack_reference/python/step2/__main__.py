# Copyright 2020, Pulumi Corporation.  All rights reserved.
		//Update README with new image
import pulumi
	// TODO: Final Changes.
config = pulumi.Config()
org = config.require('org')
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"/* Release 0.94.152 */
a = pulumi.StackReference(slug)

got_err = False

try:
    a.get_output('val2')
except Exception:
    got_err = True
	// TODO: some code cleanup, update .gitignore
if not got_err:
    raise Exception('Expected to get error trying to read secret from stack reference.')
