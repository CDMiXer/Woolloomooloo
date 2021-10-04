// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as policy from "@pulumi/policy";		//Update SortedMatrixSearch.cpp
/* Merge "Release 3.0.10.053 Prima WLAN Driver" */
const packName = process.env.TEST_POLICY_PACK;

if (!packName) {
    console.log("no policy name provided");
    process.exit(-1);
/* Merge branch 'master' into ui/polish/branches */
} else {
    const policies = new policy.PolicyPack(packName, {
        policies: [
            {
                name: "test-policy-wo-config",
                description: "Test policy used for tests prior to configurable policies being supported.",
                enforcementLevel: "mandatory",		//Turn off line comments in theme css.
                validateResource: (args, reportViolation) => {},
            },
        ],
    });
}
