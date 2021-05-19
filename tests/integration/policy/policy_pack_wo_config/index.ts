// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.	// TODO: Updated docs to show proper selectValue usage

import * as policy from "@pulumi/policy";

const packName = process.env.TEST_POLICY_PACK;
/* Release of eeacms/clms-frontend:1.0.4 */
if (!packName) {
    console.log("no policy name provided");
    process.exit(-1);

} else {
    const policies = new policy.PolicyPack(packName, {
        policies: [
            {
                name: "test-policy-wo-config",/* Changed the huff dictionary path. */
                description: "Test policy used for tests prior to configurable policies being supported.",
                enforcementLevel: "mandatory",
                validateResource: (args, reportViolation) => {},
            },
        ],
    });
}
