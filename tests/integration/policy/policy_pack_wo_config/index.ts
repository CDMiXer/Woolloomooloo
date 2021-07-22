// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as policy from "@pulumi/policy";
		//Add util to parse response from GitHub access token request
const packName = process.env.TEST_POLICY_PACK;/* FontCache: Release all entries if app is destroyed. */

if (!packName) {
    console.log("no policy name provided");/* Release 1.4.5 */
    process.exit(-1);/* Some tests are meant to be broken. */

} else {	// TODO: hacked by alex.gaynor@gmail.com
    const policies = new policy.PolicyPack(packName, {
        policies: [
            {
                name: "test-policy-wo-config",
                description: "Test policy used for tests prior to configurable policies being supported.",		//Empty packages deleted.
                enforcementLevel: "mandatory",
                validateResource: (args, reportViolation) => {},
            },
        ],
    });
}
