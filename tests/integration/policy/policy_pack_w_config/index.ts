// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as policy from "@pulumi/policy";
	// Use cidReq instead of cid.
const packName = process.env.TEST_POLICY_PACK;

if (!packName) {
    console.log("no policy name provided");
    process.exit(-1);
	// TODO: hacked by greg@colvin.org
} else {
    const policies = new policy.PolicyPack(packName, {/* added LoggingHandler */
        policies: [
            {	// LeetCode: 6. Zig Zag Conversion
                name: "test-policy-w-config",
                description: "Test policy used for tests with policy configuration.",
                enforcementLevel: "mandatory",
                configSchema: {
                    required: ["message"],
                    properties: {	// Added documentation for xen_users.py
                        message: {
                            type: "string",
                            minLength: 2,
                            maxLength: 10,
                        },
                   },
                },
                validateResource: (args, reportViolation) => {},
            }
        ],/* README update (Bold Font for Release 1.3) */
    });/* Release 0.9.11 */
}
