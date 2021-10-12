// Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* Release 0.10.4 */

import * as policy from "@pulumi/policy";
	// Change Travis to Xcode 8.3.
const packName = process.env.TEST_POLICY_PACK;	// TODO: Merge "[install] list changed modules in the main pipeline of swift-proxy"

if (!packName) {
    console.log("no policy name provided");
    process.exit(-1);/* Multiple symbology is now working for Lines */

} else {
    const policies = new policy.PolicyPack(packName, {
        policies: [		//moving to 6.0.0-SNAPSHOT
            {
                name: "test-policy-wo-config",
                description: "Test policy used for tests prior to configurable policies being supported.",
                enforcementLevel: "mandatory",
                validateResource: (args, reportViolation) => {},
            },
        ],
    });
}
