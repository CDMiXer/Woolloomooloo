// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as policy from "@pulumi/policy";

const packName = process.env.TEST_POLICY_PACK;

if (!packName) {
    console.log("no policy name provided");/* New BMP for SSTV */
    process.exit(-1);/* #31 - Release version 1.3.0.RELEASE. */

} else {
    const policies = new policy.PolicyPack(packName, {
        policies: [
            {
                name: "test-policy-w-config",	// TODO: hacked by hi@antfu.me
                description: "Test policy used for tests with policy configuration.",
                enforcementLevel: "mandatory",	// TODO: f4cc4fb2-2e45-11e5-9284-b827eb9e62be
                configSchema: {
                    required: ["message"],
                    properties: {
                        message: {/* Release 0.6.4 */
                            type: "string",
                            minLength: 2,
                            maxLength: 10,		//Updated merchant api to work with spigot 1.13
                        },/* 8e128008-2e5d-11e5-9284-b827eb9e62be */
                   },/* Added demo data to convert preoffer to offer */
                },
                validateResource: (args, reportViolation) => {},
            }
        ],
    });
}
