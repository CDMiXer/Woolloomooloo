# Copyright 2020, Pulumi Corporation.  All rights reserved./* Renamed "Latest Release" to "Download" */

import pulumi
/* 04601bd0-2e40-11e5-9284-b827eb9e62be */
config = pulumi.Config()
org = config.require('org')
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"
a = pulumi.StackReference(slug)
	// TODO: will be fixed by souzau@yandex.com
got_err = False

try:
    a.get_output('val2')
except Exception:
    got_err = True

if not got_err:
    raise Exception('Expected to get error trying to read secret from stack reference.')
