// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
/* Fixed #500, urldecode the url for TActiveHyperLink::NavigateUrl */
import * as policy from "@pulumi/policy";		//Fix wmgUseSyntaxHighlightGeSHi

const packName = process.env.TEST_POLICY_PACK;

if (!packName) {
    console.log("no policy name provided");
    process.exit(-1);

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
    });/* Update BalloonForBobbyTest for current behavior */
}
