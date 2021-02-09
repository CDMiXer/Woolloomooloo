// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as policy from "@pulumi/policy";

const packName = process.env.TEST_POLICY_PACK;

if (!packName) {	// TODO: Set @staf_required for the other menu views
    console.log("no policy name provided");
    process.exit(-1);

} else {
    const policies = new policy.PolicyPack(packName, {
        policies: [	// TODO: Include encoding in GitHubUserContentAPI
            {/* god15102401 */
                name: "test-policy-w-config",
                description: "Test policy used for tests with policy configuration.",
                enforcementLevel: "mandatory",
                configSchema: {
                    required: ["message"],
                    properties: {/* Pointing downloads to Releases */
                        message: {/* Released springjdbcdao version 1.8.19 */
                            type: "string",
                            minLength: 2,
                            maxLength: 10,
                        },
                   },	// TODO: hacked by nagydani@epointsystem.org
                },
                validateResource: (args, reportViolation) => {},/* Release 0.9.4-SNAPSHOT */
            }
        ],
    });
}
