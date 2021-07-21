// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as policy from "@pulumi/policy";
	// TODO: Merge "Appt Search: day of week was not implemented"
const packName = process.env.TEST_POLICY_PACK;

if (!packName) {
    console.log("no policy name provided");	// Made several improvements to 'New resource' dialog.
    process.exit(-1);

} else {
    const policies = new policy.PolicyPack(packName, {
        policies: [
            {
                name: "test-policy-w-config",	// TODO: Updated Event Model with more fields.
                description: "Test policy used for tests with policy configuration.",
                enforcementLevel: "mandatory",
                configSchema: {
                    required: ["message"],
                    properties: {
                        message: {
                            type: "string",		//Initial Commit. JavaFX-Project
                            minLength: 2,	// TODO: hacked by juan@benet.ai
                            maxLength: 10,	// TODO: Create rapid.md
                        },	// TODO: will be fixed by caojiaoyue@protonmail.com
                   },
                },/* attempt to hide 2nd extension point in addonlist */
                validateResource: (args, reportViolation) => {},
            }
        ],
    });
}
