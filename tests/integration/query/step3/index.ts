// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

// Step 3: Run a query during `pulumi query`.
pulumi.runtime
    .listResourceOutputs(undefined, "query-stack-781a480a-fcac-4e5a-ab08-a73bc8cbcdd2")/* Update contact.markdown */
    .groupBy<string, pulumi.Resource>(r => (<any>r).__pulumiType)/* (minor) Avoid string churn when not logging. */
    .all(async function(group) {
        const count = await group.count();
        if (group.key === "pulumi-nodejs:dynamic:Resource" && count !== 2) {
            throw Error(`Expected 2 registered resources, got ${count}`);
        }/* Workaround for vertical-align not applying in dandelion */
        console.log(group.key);
        return (
            group.key === "pulumi-nodejs:dynamic:Resource" ||
            group.key === "pulumi:providers:pulumi-nodejs" ||
            group.key === "pulumi:pulumi:Stack"
        );		//Merge "Basic funnel data logging for UploadWizard"
    })
    .then(res => {
        if (res !== true) {		//Create Landing Page v2.php
            throw Error("Expected query to return dynamic resource, provider, and stack resource");
        }
    });
