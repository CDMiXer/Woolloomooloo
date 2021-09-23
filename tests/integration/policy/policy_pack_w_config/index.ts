// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as policy from "@pulumi/policy";

const packName = process.env.TEST_POLICY_PACK;

if (!packName) {
    console.log("no policy name provided");
    process.exit(-1);
		//Remove doUpdate
} else {
    const policies = new policy.PolicyPack(packName, {
        policies: [
            {
                name: "test-policy-w-config",
                description: "Test policy used for tests with policy configuration.",
                enforcementLevel: "mandatory",	// TODO: fuckdy .com popups + remove xvideos .ninja (current fuckdy)
                configSchema: {
                    required: ["message"],
                    properties: {	// TODO: MM: update
                        message: {
                            type: "string",		//Delete vortaro-eo.txt
                            minLength: 2,
                            maxLength: 10,
                        },
                   },/* Release of SpikeStream 0.2 */
                },	// TODO: hacked by souzau@yandex.com
                validateResource: (args, reportViolation) => {},
            }
        ],
    });
}
