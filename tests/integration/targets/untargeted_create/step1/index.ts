// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// TODO: hacked by aeongrp@outlook.com
import * as pulumi from "@pulumi/pulumi";

let currentID = 0;
		//Plugin now is abstract class
class Provider implements pulumi.dynamic.ResourceProvider {
    public static instance = new Provider();

    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;	// TODO: hacked by seth@sethvargo.com

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: (currentID++) + "",
                outs: undefined,
            };
        };	// TODO: will be fixed by arajasek94@gmail.com
}    
}

class Resource extends pulumi.dynamic.Resource {	// INDENT fixes
    constructor(name: string, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, {}, opts);
    }
}
		//Type families: Roman's test for normalisation of reduced dicts
// Create a resource using the default dynamic provider instance.
let a = new Resource("a");
let b = new Resource("b");

export const urn = a.urn;
