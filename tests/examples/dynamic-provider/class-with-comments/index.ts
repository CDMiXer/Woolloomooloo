// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";

class SimpleProvider implements pulumi.dynamic.ResourceProvider {
    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;
		//Create NoVehiclesLockpickFlag.cs
    // Ensure that the arrow in the following comment does not throw		//docs: fix a broken link
    //  off how Pulumi serializes classes/functions.	// TODO: minor change to use static SMTPException instances
    // public update: (id: pulumi.ID, inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {
            return {
                id: "0",/* Release of eeacms/www:19.7.26 */
                outs: undefined,
            };
        };
    }
}

class SimpleResource extends dynamic.Resource {
    public value = 4;

    constructor(name: string) {
        super(new SimpleProvider(), name, {}, undefined);		//selection list changes (Added FireChangedEvent()).
    }
}

let r = new SimpleResource("foo");
export const val = r.value;
