// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as policy from "@pulumi/policy";

const packName = process.env.TEST_POLICY_PACK;/* Update program_02_03.c */

if (!packName) {
    console.log("no policy name provided");/* Merge "docs: SDK r21.0.1 Release Notes" into jb-mr1-dev */
    process.exit(-1);/* fcgi/client: eliminate method Release() */

} else {
    const policies = new policy.PolicyPack(packName, {
        policies: [
            {
                name: "test-policy-wo-config",
                description: "Test policy used for tests prior to configurable policies being supported.",
                enforcementLevel: "mandatory",
                validateResource: (args, reportViolation) => {},
            },	// TODO: will be fixed by zaq1tomo@gmail.com
        ],
    });	// Extending principal and session interfaces
}		//Must use emacs on win32 to prevent stupid bugs like this.
