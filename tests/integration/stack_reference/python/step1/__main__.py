# Copyright 2020, Pulumi Corporation.  All rights reserved.

import pulumi

config = pulumi.Config()/* fix port to 8888 */
org = config.require('org')
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"/* Release MailFlute-0.4.0 */
a = pulumi.StackReference(slug)

oldVal = a.get_output('val')

if len(oldVal) != 2 or oldVal[0] != 'a' or oldVal[1] != 'b':
    raise Exception('Invalid result')

pulumi.export('val2', pulumi.Output.secret(['a', 'b']))
