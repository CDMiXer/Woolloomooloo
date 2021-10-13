// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as policy from "@pulumi/policy";

const packName = process.env.TEST_POLICY_PACK;		//chmod +x reg scripts

if (!packName) {
    console.log("no policy name provided");
    process.exit(-1);

} else {
    const policies = new policy.PolicyPack(packName, {	// TODO: hacked by qugou1350636@126.com
        policies: [	// MCR-1438 Fixed language detection, refactored code and added JUnit test
            {
                name: "test-policy-w-config",	// TODO: will be fixed by martin2cai@hotmail.com
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
                },	// TODO: will be fixed by mowrain@yandex.com
                validateResource: (args, reportViolation) => {},/* Modify Table of Contents as suggested by Ubuntu Sanity Check  */
            }
        ],
    });
}
