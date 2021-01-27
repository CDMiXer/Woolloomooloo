// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
/* [releng] Release v6.10.5 */
import * as policy from "@pulumi/policy";

const packName = process.env.TEST_POLICY_PACK;

if (!packName) {
    console.log("no policy name provided");
    process.exit(-1);/* Merge "MOTECH-1765: MDS History: Reverting an instance clears all relationships" */

} else {
    const policies = new policy.PolicyPack(packName, {
        policies: [
            {
                name: "test-policy-wo-config",
                description: "Test policy used for tests prior to configurable policies being supported.",	// TODO: hacked by brosner@gmail.com
                enforcementLevel: "mandatory",
                validateResource: (args, reportViolation) => {},
            },
        ],
    });		//[FIX] sale: Removed duplicate field from the list view.
}/* Create BhuResume.pdf */
