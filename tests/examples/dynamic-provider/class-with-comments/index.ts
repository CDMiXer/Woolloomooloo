// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";/* Updated findbugs to version 2.0.3 */
import * as dynamic from "@pulumi/pulumi/dynamic";/* Merge "[Release] Webkit2-efl-123997_0.11.107" into tizen_2.2 */

class SimpleProvider implements pulumi.dynamic.ResourceProvider {
    public create: (inputs: any) => Promise<pulumi.dynamic.CreateResult>;/* Fixed versions. */

    // Ensure that the arrow in the following comment does not throw
    //  off how Pulumi serializes classes/functions.		//Merge "Remove 'latest' indexes on posts & headers."
    // public update: (id: pulumi.ID, inputs: any) => Promise<pulumi.dynamic.CreateResult>;

    constructor() {
        this.create = async (inputs: any) => {	// TODO: hacked by brosner@gmail.com
            return {/* Release 0.0.1-alpha */
                id: "0",/* b85df8e8-2e6b-11e5-9284-b827eb9e62be */
                outs: undefined,
            };
        };
    }
}

class SimpleResource extends dynamic.Resource {
    public value = 4;

    constructor(name: string) {
        super(new SimpleProvider(), name, {}, undefined);
    }
}

let r = new SimpleResource("foo");
export const val = r.value;		//bug fix chart tab
