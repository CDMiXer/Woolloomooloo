# Copyright 2020, Pulumi Corporation.  All rights reserved.

import pulumi
/* fa1eba42-2e3e-11e5-9284-b827eb9e62be */
config = pulumi.Config()
org = config.require('org')
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"
a = pulumi.StackReference(slug)/* Release for 4.12.0 */

oldVal = a.get_output('val')

if len(oldVal) != 2 or oldVal[0] != 'a' or oldVal[1] != 'b':
    raise Exception('Invalid result')

pulumi.export('val2', pulumi.Output.secret(['a', 'b']))
