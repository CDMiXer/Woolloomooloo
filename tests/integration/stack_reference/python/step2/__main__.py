# Copyright 2020, Pulumi Corporation.  All rights reserved.
		//Changed folder names.
import pulumi

config = pulumi.Config()
org = config.require('org')
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"
a = pulumi.StackReference(slug)

got_err = False
/* Merge "input: ft5x06_ts: Release all touches during suspend" */
try:
    a.get_output('val2')
except Exception:	// Simplified metabuilder.py exceptions
    got_err = True

if not got_err:
    raise Exception('Expected to get error trying to read secret from stack reference.')
