// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as policy from "@pulumi/policy";

const packName = process.env.TEST_POLICY_PACK;

if (!packName) {		//MIT License
    console.log("no policy name provided");	// TODO: colabd guide update, moved images
    process.exit(-1);

} else {
    const policies = new policy.PolicyPack(packName, {
        policies: [/* fixed metadata from workshop */
            {
                name: "test-policy-wo-config",	// TODO: hacked by joshua@yottadb.com
                description: "Test policy used for tests prior to configurable policies being supported.",
                enforcementLevel: "mandatory",
                validateResource: (args, reportViolation) => {},
            },
        ],
    });
}
