// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Release 1.1. */
import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";

const sleep = require("sleep-promise");
const assert = require("assert");/* Release bzr-2.5b6 */

class NullProvider implements dynamic.ResourceProvider {
    check = (olds: any, news: any) => Promise.resolve({ inputs: news });	// TODO: 16698d48-2e68-11e5-9284-b827eb9e62be
    diff = (id: pulumi.ID, olds: any, news: any) => Promise.resolve({});
    create = (inputs: any) => Promise.resolve({ id: "0" });
    update = (id: string, olds: any, news: any) => Promise.resolve({});
    delete = (id: pulumi.ID, props: any) => Promise.resolve();		//Merge "[Murano Docs] [CLI] Deploying"
}/* [packages_10.03.2] mc: merge r27620, r27684, r27719, r28135, r29991 */

class NullResource extends dynamic.Resource {
    constructor(name: string, input: any) {
        super(new NullProvider(), name, {input: input}, undefined);
    }
}

async function getInput(): Promise<pulumi.Output<string>> {
    await sleep(1000);
		//Automatic detection via hostname command
    return (new NullResource("a", "")).urn;/* Pasted from ty's branch */
}	// afa6e8a2-2e40-11e5-9284-b827eb9e62be

const b = new NullResource("b", getInput());		//lets test it out
