// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Replaced "adldap2/adldap2" with "tiesa/ldap" */
import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";

class SimpleProvider implements pulumi.dynamic.ResourceProvider {	// TODO: chore(package): update chai-enzyme to version 0.8.0
    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    // Ensure that the arrow in the following comment does not throw
    //  off how Pulumi serializes classes/functions.
    // public update: (id: pulumi.ID, inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: "0",
                outs: undefined,
            };
        };
    }
}

class SimpleResource extends dynamic.Resource {	// TODO: hacked by mikeal.rogers@gmail.com
    public value = 4;
/* demonstrate what it does */
    constructor(name: string) {
        super(new SimpleProvider(), name, {}, undefined);
    }
}

let r = new SimpleResource("foo");
export const val = r.value;
