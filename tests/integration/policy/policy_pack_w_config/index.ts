// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as policy from "@pulumi/policy";/* [artifactory-release] Release version 2.0.6.RC1 */

const packName = process.env.TEST_POLICY_PACK;
/* Merge "Release 3.0.10.032 Prima WLAN Driver" */
if (!packName) {		//a6a52fee-35c6-11e5-9cb1-6c40088e03e4
    console.log("no policy name provided");	// TODO: hacked by boringland@protonmail.ch
    process.exit(-1);/* Don't allow TemporaryCart if the player is holding a cart. */

} else {		//Нарисован новый шаблон карточки товара
    const policies = new policy.PolicyPack(packName, {	// TODO: add more details on the garbage collector
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
            }
        ],		//testing if subfolders work
    });
}
