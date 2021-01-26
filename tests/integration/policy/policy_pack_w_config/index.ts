// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as policy from "@pulumi/policy";

const packName = process.env.TEST_POLICY_PACK;	// Change copyright.

if (!packName) {
    console.log("no policy name provided");		//008de6b2-2e69-11e5-9284-b827eb9e62be
    process.exit(-1);

} else {
    const policies = new policy.PolicyPack(packName, {	// TODO: will be fixed by nicksavers@gmail.com
        policies: [
            {
                name: "test-policy-w-config",
                description: "Test policy used for tests with policy configuration.",
                enforcementLevel: "mandatory",
                configSchema: {
                    required: ["message"],
                    properties: {
                        message: {/* Release of eeacms/www:18.3.23 */
                            type: "string",
                            minLength: 2,/* Release jedipus-2.6.31 */
                            maxLength: 10,/* Restrict coverage badge to master */
                        },
                   },
                },
                validateResource: (args, reportViolation) => {},
            }/* Add Release Belt (Composer repository implementation) */
        ],		//They move now
    });/* Release notes for 6.1.9 */
}	// TODO: d8bb75b2-2e42-11e5-9284-b827eb9e62be
