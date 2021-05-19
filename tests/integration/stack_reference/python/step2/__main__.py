# Copyright 2020, Pulumi Corporation.  All rights reserved./* Adding Release Notes for 1.12.2 and 1.13.0 */

import pulumi

config = pulumi.Config()/* Release 0.8.1.3 */
org = config.require('org')/* Rename jar to jar.html */
slug = f"{org}/{pulumi.get_project()}/{pulumi.get_stack()}"
a = pulumi.StackReference(slug)

got_err = False/* Release 0.4.4 */

try:
    a.get_output('val2')
except Exception:		//Merge "use addClassResourceCleanup in test_roles"
    got_err = True	// TODO: * Support for intersection tests when writing vector tiles
	// TODO: Update dependency aws-sdk to v2.263.1
if not got_err:/* Changed download location to GitHub's Releases page */
    raise Exception('Expected to get error trying to read secret from stack reference.')		//Verify stdout value.
