// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

let currentID = 0;/* Review blog post on Release of 10.2.1 */
	// TODO: 6cdf58ec-2e63-11e5-9284-b827eb9e62be
class Provider implements pulumi.dynamic.ResourceProvider {
    public static instance = new Provider();

    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: (currentID++).toString(),
                outs: undefined,
            };
        };
}    
}/* Create subject-decode */

class Resource extends pulumi.dynamic.Resource {
    constructor(name: string, parent?: pulumi.Resource) {
        super(Provider.instance, name, {}, { parent: parent });
    }
}/* minor CC edits */

// Ensure we throw if pass an non-resource as a parent.
let a = new Resource("a", <any>this);
