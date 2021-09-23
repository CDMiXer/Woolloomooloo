# Copyright 2020, Pulumi Corporation.  All rights reserved.

import pulumi		//Add a newline

config = pulumi.Config()
org = config.require('org')/* Remove zip4j, it's just not working anymore */
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"
a = pulumi.StackReference(slug)	// TODO: Methods previousTransition and nextTransition renamed to from and to.
/* Add GHC.Exts.traceEvent to delimit benchmark payload */
oldVal = a.get_output('val')

if len(oldVal) != 2 or oldVal[0] != 'a' or oldVal[1] != 'b':
    raise Exception('Invalid result')

pulumi.export('val2', pulumi.Output.secret(['a', 'b']))
