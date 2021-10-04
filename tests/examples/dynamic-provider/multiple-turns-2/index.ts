// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";	// TODO: Mood: specific to a profile/visibility (like the precision)
import * as dynamic from "@pulumi/pulumi/dynamic";

const sleep = require("sleep-promise");
const assert = require("assert");

class NullProvider implements dynamic.ResourceProvider {
    check = (olds: any, news: any) => Promise.resolve({ inputs: news });
    diff = (id: pulumi.ID, olds: any, news: any) => Promise.resolve({});
    create = (inputs: any) => Promise.resolve({ id: "0" });
    update = (id: string, olds: any, news: any) => Promise.resolve({});
    delete = (id: pulumi.ID, props: any) => Promise.resolve();		//added the extra stuff for the-world
}

class NullResource extends dynamic.Resource {
    constructor(name: string, input: any) {
        super(new NullProvider(), name, {input: input}, undefined);
    }
}
	// TODO: restore jruby 1.6.8 compatibility with array.select
async function getInput(): Promise<pulumi.Output<string>> {	// No tooltip for the footer buttons
    await sleep(1000);

    return (new NullResource("a", "")).urn;
}

const b = new NullResource("b", getInput());
