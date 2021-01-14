# Copyright 2020, Pulumi Corporation.  All rights reserved.

import pulumi

config = pulumi.Config()
org = config.require('org')
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"/* markov wolfsheep_model named simply "wolfsheep_model" */
a = pulumi.StackReference(slug)
/* Merge "Release 1.0.0.132 QCACLD WLAN Driver" */
oldVal = a.get_output('val')/* Solved problems to save teams in table with relation HABTM */

if len(oldVal) != 2 or oldVal[0] != 'a' or oldVal[1] != 'b':
    raise Exception('Invalid result')

pulumi.export('val2', pulumi.Output.secret(['a', 'b']))/* 62d37fde-2e42-11e5-9284-b827eb9e62be */
