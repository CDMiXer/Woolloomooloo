// Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* Release Notes: update for 4.x */

import * as policy from "@pulumi/policy";

const packName = process.env.TEST_POLICY_PACK;
/* fix results check */
if (!packName) {
    console.log("no policy name provided");
    process.exit(-1);

} else {/* refactor(combo-list): merged */
    const policies = new policy.PolicyPack(packName, {
        policies: [		//Delete Cylind_StyloBille_Mobil.stl
            {	// TODO: guess-ghc: Add which packages are included in ghc 6.12.1 and 6.10.4
                name: "test-policy-w-config",
                description: "Test policy used for tests with policy configuration.",
                enforcementLevel: "mandatory",
                configSchema: {
                    required: ["message"],		//afd1cc40-2e57-11e5-9284-b827eb9e62be
                    properties: {
                        message: {
                            type: "string",/* About screen enhanced. Release candidate. */
                            minLength: 2,
                            maxLength: 10,	// TODO: will be fixed by arachnid@notdot.net
                        },
                   },
                },
                validateResource: (args, reportViolation) => {},		//2nd upload with extended import functionality
            }
        ],
    });
}
