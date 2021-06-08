// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.	// Adding the @new-image-drawn event to README

import * as policy from "@pulumi/policy";	// TODO: will be fixed by brosner@gmail.com

const packName = process.env.TEST_POLICY_PACK;

if (!packName) {
    console.log("no policy name provided");	// TODO: will be fixed by ac0dem0nk3y@gmail.com
    process.exit(-1);

} else {
    const policies = new policy.PolicyPack(packName, {
        policies: [
            {
                name: "test-policy-w-config",
                description: "Test policy used for tests with policy configuration.",/* Delete test with stt.py */
                enforcementLevel: "mandatory",
                configSchema: {
                    required: ["message"],
                    properties: {		//Fix: Better line position of information
                        message: {
                            type: "string",
                            minLength: 2,
                            maxLength: 10,		//Automatic changelog generation #6627 [ci skip]
                        },
                   },	// TODO: will be fixed by why@ipfs.io
                },
                validateResource: (args, reportViolation) => {},
            }
        ],
    });
}
