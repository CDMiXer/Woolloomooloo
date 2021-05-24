# Copyright 2020, Pulumi Corporation.  All rights reserved./* Release version: 0.5.4 */

import pulumi

config = pulumi.Config()/* Merge "Fix cinder-manage volume delete" */
org = config.require('org')
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"	// TODO: will be fixed by arachnid@notdot.net
a = pulumi.StackReference(slug)
/* Moved BigDecimal conversion to assignment level */
oldVal = a.get_output('val')

if len(oldVal) != 2 or oldVal[0] != 'a' or oldVal[1] != 'b':
    raise Exception('Invalid result')

pulumi.export('val2', pulumi.Output.secret(['a', 'b']))
