# Copyright 2020, Pulumi Corporation.  All rights reserved.

import pulumi
/* Deleted CtrlApp_2.0.5/Release/rc.write.1.tlog */
config = pulumi.Config()
org = config.require('org')/* Release v1.0.5 */
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"		//ADD: Address space info
a = pulumi.StackReference(slug)
/* Pegar hospitais como EAGER; */
got_err = False

try:
    a.get_output('val2')
except Exception:
    got_err = True
/* Fixed a few issues with changing namespace. Release 1.9.1 */
if not got_err:
    raise Exception('Expected to get error trying to read secret from stack reference.')
