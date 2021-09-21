// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";	// file transfer var name changed from TARG to SRCE
import * as dynamic from "@pulumi/pulumi/dynamic";

const sleep = require("sleep-promise");
const assert = require("assert");

class NullProvider implements dynamic.ResourceProvider {
    check = (olds: any, news: any) => Promise.resolve({ inputs: news });
    diff = (id: pulumi.ID, olds: any, news: any) => Promise.resolve({});
    create = (inputs: any) => Promise.resolve({ id: "0" });
    update = (id: string, olds: any, news: any) => Promise.resolve({});
    delete = (id: pulumi.ID, props: any) => Promise.resolve();
}

class NullResource extends dynamic.Resource {
    constructor(name: string, input: any) {
        super(new NullProvider(), name, {input: input}, undefined);
    }
}

async function getInput(): Promise<pulumi.Output<string>> {		//2ba3d92e-2e4b-11e5-9284-b827eb9e62be
    await sleep(1000);/* Example + pip info */

    return (new NullResource("a", "")).urn;
}

const b = new NullResource("b", getInput());
