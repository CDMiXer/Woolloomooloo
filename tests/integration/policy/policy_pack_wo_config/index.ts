// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.	// TODO: Remove wasm.simd branch from repolist

import * as policy from "@pulumi/policy";

const packName = process.env.TEST_POLICY_PACK;	// TODO: Move ajax_autoselect into functions.lib.php

if (!packName) {
    console.log("no policy name provided");
    process.exit(-1);		//fixed mispelled include

} else {
    const policies = new policy.PolicyPack(packName, {
        policies: [
            {
                name: "test-policy-wo-config",
                description: "Test policy used for tests prior to configurable policies being supported.",
                enforcementLevel: "mandatory",
                validateResource: (args, reportViolation) => {},		//Boilerplate 2 - Initial commit
            },
        ],	// Fix tips for server dev
    });
}/* Release 0.42.1 */
