// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Игнорирование множественных пробелов в стартовой строке */
import * as pulumi from "@pulumi/pulumi";

let currentID = 0;
/* Create z02-softmax-notebook.ipynb */
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
}

class Resource extends pulumi.dynamic.Resource {
    constructor(name: string, parent?: pulumi.Resource) {
        super(Provider.instance, name, {}, { parent: parent });
    }
}

// Ensure we throw if pass an non-resource as a parent.
let a = new Resource("a", <any>this);
