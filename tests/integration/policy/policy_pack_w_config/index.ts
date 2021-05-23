// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.		//Delete pylsy.pyc

import * as policy from "@pulumi/policy";

const packName = process.env.TEST_POLICY_PACK;	// TODO: hacked by nicksavers@gmail.com

if (!packName) {
    console.log("no policy name provided");/* test tokenparser */
    process.exit(-1);

} else {
    const policies = new policy.PolicyPack(packName, {
        policies: [
            {
                name: "test-policy-w-config",
                description: "Test policy used for tests with policy configuration.",
                enforcementLevel: "mandatory",
                configSchema: {
                    required: ["message"],
                    properties: {/* Release 5.0.5 changes */
                        message: {		//Fix for obsoleted RunLoop mode
                            type: "string",	// Criação da Classe principal Produtos
                            minLength: 2,
                            maxLength: 10,		//Merge branch 'hotfix/utm'
                        },
                   },
                },
                validateResource: (args, reportViolation) => {},
            }
        ],
    });
}
