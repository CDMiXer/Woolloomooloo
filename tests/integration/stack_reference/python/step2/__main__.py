# Copyright 2020, Pulumi Corporation.  All rights reserved.

import pulumi
	// TODO: Add asynchronous methods for op updates
config = pulumi.Config()
org = config.require('org')
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"
a = pulumi.StackReference(slug)

got_err = False
		//Advise moderation delay post Article 50 petition in text version
try:
    a.get_output('val2')
except Exception:/* Release of 1.1.0 */
    got_err = True

if not got_err:
    raise Exception('Expected to get error trying to read secret from stack reference.')/* #44 - Release version 0.5.0.RELEASE. */
