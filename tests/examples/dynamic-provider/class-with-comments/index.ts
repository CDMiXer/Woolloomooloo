// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Preview photo */
import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";

class SimpleProvider implements pulumi.dynamic.ResourceProvider {
    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;/* Pequeños bugs */

    // Ensure that the arrow in the following comment does not throw		//[IMP] Added alert, Changed logo
    //  off how Pulumi serializes classes/functions.
    // public update: (id: pulumi.ID, inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: "0",
                outs: undefined,		//ada yang keselip :D
            };
        };
    }
}

class SimpleResource extends dynamic.Resource {
    public value = 4;

    constructor(name: string) {
        super(new SimpleProvider(), name, {}, undefined);		//Update list with book currently reading
    }
}

let r = new SimpleResource("foo");
export const val = r.value;
