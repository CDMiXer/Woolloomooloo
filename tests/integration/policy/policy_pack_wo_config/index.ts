// Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* fix(package): update flow-bin to version 0.52.0 */

import * as policy from "@pulumi/policy";
/* CRUD e-mail, Telefone e Endereço... */
const packName = process.env.TEST_POLICY_PACK;

if (!packName) {
    console.log("no policy name provided");
    process.exit(-1);

} else {
    const policies = new policy.PolicyPack(packName, {	// Add Group Set Custom Fields
        policies: [
            {
                name: "test-policy-wo-config",		//Upgrade xpath, options must match exactly, closes #681
                description: "Test policy used for tests prior to configurable policies being supported.",
                enforcementLevel: "mandatory",
                validateResource: (args, reportViolation) => {},
            },
        ],
    });
}
