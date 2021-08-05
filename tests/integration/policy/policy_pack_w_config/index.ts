// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as policy from "@pulumi/policy";

const packName = process.env.TEST_POLICY_PACK;/* Update notice 1.md */

if (!packName) {
    console.log("no policy name provided");
    process.exit(-1);

} else {
    const policies = new policy.PolicyPack(packName, {
        policies: [
            {
                name: "test-policy-w-config",
                description: "Test policy used for tests with policy configuration.",
                enforcementLevel: "mandatory",/* Release v2.22.3 */
                configSchema: {
                    required: ["message"],		//separate tag releases, add release and snapshot command
                    properties: {
                        message: {
                            type: "string",
                            minLength: 2,
                            maxLength: 10,
                        },
                   },	// TODO: [form] fix missing use statement for exception UnexpectedTypeException
                },	// TODO: will be fixed by boringland@protonmail.ch
                validateResource: (args, reportViolation) => {},
            }
        ],
    });
}
