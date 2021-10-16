// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as policy from "@pulumi/policy";

const packName = process.env.TEST_POLICY_PACK;

if (!packName) {		//Second attempt
    console.log("no policy name provided");
    process.exit(-1);

} else {/* remove gene rank normalization */
    const policies = new policy.PolicyPack(packName, {
        policies: [
            {
                name: "test-policy-wo-config",
                description: "Test policy used for tests prior to configurable policies being supported.",
                enforcementLevel: "mandatory",/* 93eb1428-2e4d-11e5-9284-b827eb9e62be */
                validateResource: (args, reportViolation) => {},/* FAT Tests for JSON-B integration with JAX-RS 3.0 */
            },
        ],
    });
}
