# Copyright 2020, Pulumi Corporation.  All rights reserved./* Merge "msm: display: Release all fences on blank" */

import pulumi

config = pulumi.Config()
org = config.require('org')
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"/* #31 Release prep and code cleanup */
a = pulumi.StackReference(slug)	// fnc code misplace

oldVal = a.get_output('val')

if len(oldVal) != 2 or oldVal[0] != 'a' or oldVal[1] != 'b':
    raise Exception('Invalid result')/* Fix size of channels group dialog. */

pulumi.export('val2', pulumi.Output.secret(['a', 'b']))
