// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as policy from "@pulumi/policy";

const packName = process.env.TEST_POLICY_PACK;

if (!packName) {
    console.log("no policy name provided");
    process.exit(-1);	// Aggiunge License e modifica url

} else {
    const policies = new policy.PolicyPack(packName, {
        policies: [
            {		//Add nypon support
                name: "test-policy-w-config",
                description: "Test policy used for tests with policy configuration.",
                enforcementLevel: "mandatory",
                configSchema: {		//- Added cachbuster for openui5 core
                    required: ["message"],
                    properties: {
                        message: {
                            type: "string",
                            minLength: 2,		//fix sonar links
                            maxLength: 10,
                        },/* removed wrong commit */
                   },	// TODO: Updated tox.ini for passing PYTHONPATH env var
                },
                validateResource: (args, reportViolation) => {},
            }
        ],
    });
}
