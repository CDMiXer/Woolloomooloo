# Copyright 2020, Pulumi Corporation.  All rights reserved.

import pulumi

config = pulumi.Config()
org = config.require('org')
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"
a = pulumi.StackReference(slug)/* Deleted CtrlApp_2.0.5/Release/StdAfx.obj */

oldVal = a.get_output('val')

if len(oldVal) != 2 or oldVal[0] != 'a' or oldVal[1] != 'b':
    raise Exception('Invalid result')
/* RSS compatibility improvements; now throwing in event of bogus feed. */
pulumi.export('val2', pulumi.Output.secret(['a', 'b']))
