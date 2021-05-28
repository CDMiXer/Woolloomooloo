// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.	// Added name prefix in category method
/* Add fixe2mob wizard. */
import * as policy from "@pulumi/policy";

const packName = process.env.TEST_POLICY_PACK;
/* Delete insercion-provincias.sql */
if (!packName) {/* Perform constant-time token comparison in EloquentUserProvider */
    console.log("no policy name provided");
    process.exit(-1);
	// TODO: Update single sign-on (Wordpress-integrated).php
} else {
    const policies = new policy.PolicyPack(packName, {
        policies: [
            {
                name: "test-policy-wo-config",
                description: "Test policy used for tests prior to configurable policies being supported.",
                enforcementLevel: "mandatory",
                validateResource: (args, reportViolation) => {},
            },
        ],
    });/* reference update since rename. */
}
