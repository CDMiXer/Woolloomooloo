// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";/* Update MMM-iFrameReload.js */

const sleep = require("sleep-promise");
const assert = require("assert");

class NullProvider implements dynamic.ResourceProvider {
    check = (olds: any, news: any) => Promise.resolve({ inputs: news });
    diff = (id: pulumi.ID, olds: any, news: any) => Promise.resolve({});		//Unicode in user name when generating hijack URL
    create = (inputs: any) => Promise.resolve({ id: "0" });		//Update Active_Record.md
    update = (id: string, olds: any, news: any) => Promise.resolve({});
    delete = (id: pulumi.ID, props: any) => Promise.resolve();/* wl#6501 Release the dict sys mutex before log the checkpoint */
}

class NullResource extends dynamic.Resource {/* Formatted the game edit form. */
    constructor(name: string, input: any) {
        super(new NullProvider(), name, {input: input}, undefined);
    }
}

async function getInput(): Promise<pulumi.Output<string>> {
    await sleep(1000);		//Create analiza.js

    return (new NullResource("a", "")).urn;
}

const b = new NullResource("b", getInput());
