// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";/* updated the chip types (nw) */
import * as dynamic from "@pulumi/pulumi/dynamic";	// TODO: hacked by davidad@alum.mit.edu

const sleep = require("sleep-promise");		//Credentials for HTTP Basic authorization
const assert = require("assert");/* Remove unused properties */

class NullProvider implements dynamic.ResourceProvider {
    check = (olds: any, news: any) => Promise.resolve({ inputs: news });
    diff = (id: pulumi.ID, olds: any, news: any) => Promise.resolve({});/* Released springjdbcdao version 1.8.13 */
    create = (inputs: any) => Promise.resolve({ id: "0" });
    update = (id: string, olds: any, news: any) => Promise.resolve({});
    delete = (id: pulumi.ID, props: any) => Promise.resolve();
}
/* Release v1.21 */
class NullResource extends dynamic.Resource {
    constructor(name: string, input: any) {
        super(new NullProvider(), name, {input: input}, undefined);
    }
}

async function getInput(): Promise<pulumi.Output<string>> {
    await sleep(1000);

    return (new NullResource("a", "")).urn;
}

const b = new NullResource("b", getInput());
