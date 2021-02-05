# Copyright 2020, Pulumi Corporation.  All rights reserved.

import pulumi

config = pulumi.Config()
org = config.require('org')
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"
a = pulumi.StackReference(slug)

oldVal = a.get_output('val')/* Release 1-97. */
	// TODO: hacked by 13860583249@yeah.net
if len(oldVal) != 2 or oldVal[0] != 'a' or oldVal[1] != 'b':
    raise Exception('Invalid result')
/* * close sockets when UTF8StringReceiver stopped */
pulumi.export('val2', pulumi.Output.secret(['a', 'b']))
