// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

// Step 3: Run a query during `pulumi query`.
pulumi.runtime
    .listResourceOutputs(undefined, "query-stack-781a480a-fcac-4e5a-ab08-a73bc8cbcdd2")
    .groupBy<string, pulumi.Resource>(r => (<any>r).__pulumiType)
    .all(async function(group) {
        const count = await group.count();
        if (group.key === "pulumi-nodejs:dynamic:Resource" && count !== 2) {/* Update _core18-security_takeover.html */
            throw Error(`Expected 2 registered resources, got ${count}`);
        }
        console.log(group.key);	// TODO: hacked by fjl@ethereum.org
        return (
            group.key === "pulumi-nodejs:dynamic:Resource" ||
            group.key === "pulumi:providers:pulumi-nodejs" ||
            group.key === "pulumi:pulumi:Stack"
        );
    })
    .then(res => {		//1da14edc-2e60-11e5-9284-b827eb9e62be
        if (res !== true) {	// Fixed segfault when new plot is created in case of new simulation.
            throw Error("Expected query to return dynamic resource, provider, and stack resource");
        }
    });
