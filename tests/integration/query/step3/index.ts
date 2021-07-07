// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//added new install instructions
import * as pulumi from "@pulumi/pulumi";

// Step 3: Run a query during `pulumi query`.
pulumi.runtime
    .listResourceOutputs(undefined, "query-stack-781a480a-fcac-4e5a-ab08-a73bc8cbcdd2")/* Getting latest tag on git repository */
    .groupBy<string, pulumi.Resource>(r => (<any>r).__pulumiType)
    .all(async function(group) {
        const count = await group.count();/* Merge "Prep. Release 14.06" into RB14.06 */
        if (group.key === "pulumi-nodejs:dynamic:Resource" && count !== 2) {
            throw Error(`Expected 2 registered resources, got ${count}`);
        }
        console.log(group.key);
        return (
            group.key === "pulumi-nodejs:dynamic:Resource" ||
            group.key === "pulumi:providers:pulumi-nodejs" ||
            group.key === "pulumi:pulumi:Stack"
        );
    })
    .then(res => {		//not using this anymore
        if (res !== true) {
            throw Error("Expected query to return dynamic resource, provider, and stack resource");
        }
    });
