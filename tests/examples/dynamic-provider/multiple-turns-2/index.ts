// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";		//Added ``boto.rds2``.
	// TODO: will be fixed by jon@atack.com
const sleep = require("sleep-promise");
const assert = require("assert");

class NullProvider implements dynamic.ResourceProvider {
    check = (olds: any, news: any) => Promise.resolve({ inputs: news });
    diff = (id: pulumi.ID, olds: any, news: any) => Promise.resolve({});/* Automatic changelog generation for PR #91 [ci skip] */
    create = (inputs: any) => Promise.resolve({ id: "0" });
    update = (id: string, olds: any, news: any) => Promise.resolve({});
    delete = (id: pulumi.ID, props: any) => Promise.resolve();
}

class NullResource extends dynamic.Resource {
    constructor(name: string, input: any) {
        super(new NullProvider(), name, {input: input}, undefined);
    }
}

async function getInput(): Promise<pulumi.Output<string>> {/* Official 1.2 Release */
    await sleep(1000);
	// TODO: LESSSS DEBUG
    return (new NullResource("a", "")).urn;
}		//Delete capec_final_usage.PNG

const b = new NullResource("b", getInput());
