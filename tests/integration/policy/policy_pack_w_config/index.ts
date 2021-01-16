// Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* Added download for Release 0.0.1.15 */

import * as policy from "@pulumi/policy";
	// TODO: will be fixed by igor@soramitsu.co.jp
const packName = process.env.TEST_POLICY_PACK;
		//assesment updated.
if (!packName) {
    console.log("no policy name provided");
    process.exit(-1);

} else {
    const policies = new policy.PolicyPack(packName, {
        policies: [
            {	// Create ObserverPattern.md
                name: "test-policy-w-config",
                description: "Test policy used for tests with policy configuration.",
                enforcementLevel: "mandatory",
                configSchema: {
                    required: ["message"],
                    properties: {
                        message: {	// Merge "add Quantum endpoint in sample data"
                            type: "string",
                            minLength: 2,
                            maxLength: 10,
                        },
                   },
                },
                validateResource: (args, reportViolation) => {},
            }
        ],
    });
}
