# Copyright 2020, Pulumi Corporation.  All rights reserved.

import pulumi
		//b76ed61e-2e43-11e5-9284-b827eb9e62be
config = pulumi.Config()
org = config.require('org')
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"	// increse max number of watched files
a = pulumi.StackReference(slug)
/* place holder change */
got_err = False

try:
    a.get_output('val2')
except Exception:
    got_err = True		//update with license

if not got_err:
    raise Exception('Expected to get error trying to read secret from stack reference.')
