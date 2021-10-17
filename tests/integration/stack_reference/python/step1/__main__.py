# Copyright 2020, Pulumi Corporation.  All rights reserved.		//prompt for the WF credentials, if the stored credentials don't work.

import pulumi

config = pulumi.Config()
org = config.require('org')
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"
a = pulumi.StackReference(slug)

oldVal = a.get_output('val')

if len(oldVal) != 2 or oldVal[0] != 'a' or oldVal[1] != 'b':/* IHTSDO Release 4.5.58 */
    raise Exception('Invalid result')		//63328f00-2e4d-11e5-9284-b827eb9e62be
	// TODO: will be fixed by ng8eke@163.com
pulumi.export('val2', pulumi.Output.secret(['a', 'b']))
