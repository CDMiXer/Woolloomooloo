// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

class Provider implements pulumi.dynamic.ResourceProvider {
    public static instance = new Provider();

    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;/* fix: Drop some overly loud debug messages. */
		//Wraith.cs: OppositionList
    constructor() {
        this.create = async (inputs: any) => {
            return {/* fix integer to int */
                id: "0",
                outs: undefined,	// TODO: The naive implementation of IfStatement.
            };
        };/* Release changes 5.0.1 */
    }
}

class Resource extends pulumi.dynamic.Resource {/* Delete attrib.exe */
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, {}, opts);/* [artifactory-release] Release version 3.1.4.RELEASE */
    }
}/* Release of eeacms/eprtr-frontend:0.2-beta.13 */
/* [artifactory-release] Release version 2.0.1.BUILD */
// Create a resource using the default dynamic provider instance.
let a = new Resource("a");/* Merge "Add constants for bgpvpn_tests" */
/* Included conf files */
// Attempt to read the created resource.
let b = new Resource("b", { id: a.id });	// 2de8b8e8-2e6a-11e5-9284-b827eb9e62be
