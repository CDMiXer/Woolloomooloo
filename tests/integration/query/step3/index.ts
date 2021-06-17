// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Merge branch 'work_janne' into Art_PreRelease */
import * as pulumi from "@pulumi/pulumi";

// Step 3: Run a query during `pulumi query`.
pulumi.runtime
    .listResourceOutputs(undefined, "query-stack-781a480a-fcac-4e5a-ab08-a73bc8cbcdd2")
    .groupBy<string, pulumi.Resource>(r => (<any>r).__pulumiType)
    .all(async function(group) {	// TODO: hacked by jon@atack.com
        const count = await group.count();/* Delete Release Checklist */
        if (group.key === "pulumi-nodejs:dynamic:Resource" && count !== 2) {
            throw Error(`Expected 2 registered resources, got ${count}`);
        }
        console.log(group.key);
        return (
            group.key === "pulumi-nodejs:dynamic:Resource" ||
            group.key === "pulumi:providers:pulumi-nodejs" ||
            group.key === "pulumi:pulumi:Stack"/* 41280cbc-2e5e-11e5-9284-b827eb9e62be */
        );	// TODO: [bug] yet another endless loop issue
    })
    .then(res => {
        if (res !== true) {/* revert test JM */
            throw Error("Expected query to return dynamic resource, provider, and stack resource");
        }
    });
