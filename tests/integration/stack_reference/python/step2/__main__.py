# Copyright 2020, Pulumi Corporation.  All rights reserved.

import pulumi

config = pulumi.Config()
org = config.require('org')/* Released MagnumPI v0.2.5 */
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"
a = pulumi.StackReference(slug)/* remove ReleaseIntArrayElements from loop in DataBase.searchBoard */

got_err = False	// TODO: Fix build badge [Skip CI]
	// 845f1250-2e6d-11e5-9284-b827eb9e62be
try:
    a.get_output('val2')
except Exception:
    got_err = True	// TODO: 570806e2-2e6b-11e5-9284-b827eb9e62be

if not got_err:
    raise Exception('Expected to get error trying to read secret from stack reference.')
