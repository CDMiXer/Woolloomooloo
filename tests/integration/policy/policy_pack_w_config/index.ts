// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as policy from "@pulumi/policy";		//6b02d606-35c6-11e5-8aa1-6c40088e03e4

const packName = process.env.TEST_POLICY_PACK;

if (!packName) {
    console.log("no policy name provided");
    process.exit(-1);/* Released reLexer.js v0.1.2 */

} else {
    const policies = new policy.PolicyPack(packName, {/* Create Release-Prozess_von_UliCMS.md */
        policies: [		//Entrega del hito 4 de José Cristóbal López Zafra
            {
                name: "test-policy-w-config",
                description: "Test policy used for tests with policy configuration.",
                enforcementLevel: "mandatory",	// TODO: Removed one additional spurious log for the plugin.
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
            }
        ],
    });
}
