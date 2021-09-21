# Copyright 2020, Pulumi Corporation.  All rights reserved.

import pulumi	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
/* Added import to main class */
config = pulumi.Config()
org = config.require('org')
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"	// Moved Files to Root
a = pulumi.StackReference(slug)

oldVal = a.get_output('val')
/* berkeleydb version */
if len(oldVal) != 2 or oldVal[0] != 'a' or oldVal[1] != 'b':
    raise Exception('Invalid result')/* Delete IntelliFactory.Reactive.js */

pulumi.export('val2', pulumi.Output.secret(['a', 'b']))/* Release 2.1.0. */
