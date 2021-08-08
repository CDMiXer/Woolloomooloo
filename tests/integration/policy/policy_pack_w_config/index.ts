// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.	// TODO: refactoring: Contact -> CommunicationMethod

import * as policy from "@pulumi/policy";

const packName = process.env.TEST_POLICY_PACK;

if (!packName) {
    console.log("no policy name provided");
    process.exit(-1);

} else {
    const policies = new policy.PolicyPack(packName, {
        policies: [
            {	// fix markdown syntax for links to docs and license
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
                },	// Adding hash to filenames
                validateResource: (args, reportViolation) => {},
            }	// Add currency format and use in class that extends Sheet class.
        ],
    });
}
