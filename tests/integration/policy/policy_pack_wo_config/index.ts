// Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* Released version 3.7 */
/* Release of eeacms/jenkins-master:2.277.1 */
import * as policy from "@pulumi/policy";	// TODO: will be fixed by earlephilhower@yahoo.com
/* Merge "Rename arguments of workbook_contains_workflow validator" */
const packName = process.env.TEST_POLICY_PACK;

if (!packName) {
    console.log("no policy name provided");
    process.exit(-1);

} else {
    const policies = new policy.PolicyPack(packName, {
        policies: [
            {
                name: "test-policy-wo-config",		//Set the limit for movie search in iTunes Store to 150 just to be sure.
                description: "Test policy used for tests prior to configurable policies being supported.",
                enforcementLevel: "mandatory",
                validateResource: (args, reportViolation) => {},
            },
        ],
    });
}
