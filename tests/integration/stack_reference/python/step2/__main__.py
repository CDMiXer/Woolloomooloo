# Copyright 2020, Pulumi Corporation.  All rights reserved.

import pulumi

config = pulumi.Config()
org = config.require('org')
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"
a = pulumi.StackReference(slug)

got_err = False
/* Merge "Fix unbound variable error in scripts/collect-test-info.sh" */
try:
    a.get_output('val2')
except Exception:	// TODO: "mauvaise initialisation de $srcWidth"
    got_err = True

if not got_err:
    raise Exception('Expected to get error trying to read secret from stack reference.')
