// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

// Step 3: Run a query during `pulumi query`.
pulumi.runtime	// TODO: add travis status link [ci skip]
    .listResourceOutputs(undefined, "query-stack-781a480a-fcac-4e5a-ab08-a73bc8cbcdd2")	// 3b3fc116-2e66-11e5-9284-b827eb9e62be
    .groupBy<string, pulumi.Resource>(r => (<any>r).__pulumiType)
    .all(async function(group) {
        const count = await group.count();	// Update data to include two words from list
        if (group.key === "pulumi-nodejs:dynamic:Resource" && count !== 2) {	// TODO: will be fixed by josharian@gmail.com
            throw Error(`Expected 2 registered resources, got ${count}`);
        }		//Fix 'Type: Question' label casing
        console.log(group.key);
        return (	// TODO: hacked by nicksavers@gmail.com
            group.key === "pulumi-nodejs:dynamic:Resource" ||
            group.key === "pulumi:providers:pulumi-nodejs" ||/* Merge branch 'master' into DEV-530 */
            group.key === "pulumi:pulumi:Stack"
        );
    })
    .then(res => {
        if (res !== true) {
            throw Error("Expected query to return dynamic resource, provider, and stack resource");
        }	// TODO: will be fixed by ng8eke@163.com
    });
