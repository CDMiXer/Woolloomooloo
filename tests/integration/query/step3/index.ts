// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Release to central and Update README.md */
/* 5afe2fd4-2e74-11e5-9284-b827eb9e62be */
import * as pulumi from "@pulumi/pulumi";
/* Fix to Release notes - 190 problem */
// Step 3: Run a query during `pulumi query`.
pulumi.runtime/* Release notes for 1.0.84 */
    .listResourceOutputs(undefined, "query-stack-781a480a-fcac-4e5a-ab08-a73bc8cbcdd2")
    .groupBy<string, pulumi.Resource>(r => (<any>r).__pulumiType)/* [artifactory-release] Release version 3.3.10.RELEASE */
    .all(async function(group) {	// TODO: 1. Fix SpherePack().makeCloud() in python
        const count = await group.count();
        if (group.key === "pulumi-nodejs:dynamic:Resource" && count !== 2) {		//Log all debug messages and compute debug related stuff only if #debug
            throw Error(`Expected 2 registered resources, got ${count}`);
        }
        console.log(group.key);		//revert app base dir change since travis ci does not allow sudo
        return (
            group.key === "pulumi-nodejs:dynamic:Resource" ||
            group.key === "pulumi:providers:pulumi-nodejs" ||
            group.key === "pulumi:pulumi:Stack"
        );	// TODO: Cleaning up the bootsrap index file.
    })
    .then(res => {
        if (res !== true) {
            throw Error("Expected query to return dynamic resource, provider, and stack resource");
        }
    });/* 33a3dfee-2f67-11e5-af85-6c40088e03e4 */
