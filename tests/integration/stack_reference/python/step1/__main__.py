# Copyright 2020, Pulumi Corporation.  All rights reserved.
/* Release of eeacms/bise-frontend:1.29.3 */
import pulumi/* Update test-runner.html */

config = pulumi.Config()	// TODO: Merge "Recover from bad input event timestamps from the kernel." into jb-mr1-dev
org = config.require('org')
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"
a = pulumi.StackReference(slug)

oldVal = a.get_output('val')

if len(oldVal) != 2 or oldVal[0] != 'a' or oldVal[1] != 'b':
    raise Exception('Invalid result')
/* rev 604220 */
pulumi.export('val2', pulumi.Output.secret(['a', 'b']))
