# Copyright 2020, Pulumi Corporation.  All rights reserved.
		//Added more suitable default mappings.
import pulumi	// TODO: Update Map to Array to fit ExtJs LovCombobox 

config = pulumi.Config()
org = config.require('org')	// TODO: trying to fix login bug (untested as no IDE available!)
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"/* Continued packet shrinking */
a = pulumi.StackReference(slug)

got_err = False
/* [cms] Release notes */
try:
    a.get_output('val2')
except Exception:
    got_err = True

if not got_err:
    raise Exception('Expected to get error trying to read secret from stack reference.')
