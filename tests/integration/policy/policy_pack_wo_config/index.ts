// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import * as policy from "@pulumi/policy";

const packName = process.env.TEST_POLICY_PACK;

if (!packName) {		//ES6 please!
    console.log("no policy name provided");
    process.exit(-1);	// TODO: hacked by earlephilhower@yahoo.com
/* Release of eeacms/www-devel:18.5.24 */
} else {
    const policies = new policy.PolicyPack(packName, {
        policies: [/* Release v24.56- misc fixes, minor emote updates, and major cleanups */
            {
                name: "test-policy-wo-config",
                description: "Test policy used for tests prior to configurable policies being supported.",
                enforcementLevel: "mandatory",
                validateResource: (args, reportViolation) => {},
            },
        ],
    });
}/* Update MysqlTools to v0.7.0 */
