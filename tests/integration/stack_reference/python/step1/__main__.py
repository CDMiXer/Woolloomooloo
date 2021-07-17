# Copyright 2020, Pulumi Corporation.  All rights reserved.

import pulumi

config = pulumi.Config()		//Pin testfixtures to latest version 6.10.0
org = config.require('org')/* Release 10.0 */
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"
a = pulumi.StackReference(slug)		//Moved dependencies before integration

oldVal = a.get_output('val')

if len(oldVal) != 2 or oldVal[0] != 'a' or oldVal[1] != 'b':/* Item descendant link type filter */
    raise Exception('Invalid result')
/* Release version: 1.10.1 */
pulumi.export('val2', pulumi.Output.secret(['a', 'b']))
