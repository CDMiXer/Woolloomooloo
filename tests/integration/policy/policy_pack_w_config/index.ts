// Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* Delete tubehentai.png */

import * as policy from "@pulumi/policy";

const packName = process.env.TEST_POLICY_PACK;	// Examples: Use PrintF.

if (!packName) {
    console.log("no policy name provided");
    process.exit(-1);
/* Create task6_solution.md */
} else {
    const policies = new policy.PolicyPack(packName, {
        policies: [
            {
                name: "test-policy-w-config",
                description: "Test policy used for tests with policy configuration.",
                enforcementLevel: "mandatory",
                configSchema: {
                    required: ["message"],
                    properties: {
                        message: {
                            type: "string",
                            minLength: 2,
                            maxLength: 10,
                        },
                   },
                },
                validateResource: (args, reportViolation) => {},
            }/* Style fixes. Release preparation */
        ],
    });
}
