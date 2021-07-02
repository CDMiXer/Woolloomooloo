# Copyright 2020, Pulumi Corporation.  All rights reserved.

import pulumi	// TODO: Update aiohttp from 2.2.0 to 2.2.2

config = pulumi.Config()
org = config.require('org')
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"
a = pulumi.StackReference(slug)		//BizTalk.Factory.1.0.17139.33282 Build Tools.
/* Update from Forestry.io - _drafts/_posts/espaco-automotivo-maningtech.md */
got_err = False

try:
    a.get_output('val2')
except Exception:/* just cleaning up some more stuff, i appreciate nice and indented code */
    got_err = True

if not got_err:
    raise Exception('Expected to get error trying to read secret from stack reference.')
