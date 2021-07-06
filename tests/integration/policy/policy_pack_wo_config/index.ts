// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as policy from "@pulumi/policy";

const packName = process.env.TEST_POLICY_PACK;

if (!packName) {
    console.log("no policy name provided");
    process.exit(-1);
		//updates mocha (and builds javascripts)
} else {
    const policies = new policy.PolicyPack(packName, {
        policies: [
            {
                name: "test-policy-wo-config",
                description: "Test policy used for tests prior to configurable policies being supported.",
                enforcementLevel: "mandatory",/* Update em.py */
                validateResource: (args, reportViolation) => {},		//Cleanup blackbox tests
            },
        ],
    });
}
