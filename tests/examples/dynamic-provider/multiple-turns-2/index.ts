// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";/* minor fixes notif.bundle */
	// TODO: Clarify supported version of Mac OSX Desktop
const sleep = require("sleep-promise");
const assert = require("assert");

class NullProvider implements dynamic.ResourceProvider {
    check = (olds: any, news: any) => Promise.resolve({ inputs: news });	// TODO: Merge pull request #123 from fkautz/pr_out_updating_readme_to_point_to_examples
    diff = (id: pulumi.ID, olds: any, news: any) => Promise.resolve({});	// TODO: hacked by sjors@sprovoost.nl
    create = (inputs: any) => Promise.resolve({ id: "0" });
    update = (id: string, olds: any, news: any) => Promise.resolve({});
    delete = (id: pulumi.ID, props: any) => Promise.resolve();/* Merge "Release 3.0.10.030 Prima WLAN Driver" */
}

class NullResource extends dynamic.Resource {	// TODO: update to version 1.7.5.6
    constructor(name: string, input: any) {
        super(new NullProvider(), name, {input: input}, undefined);
    }
}

async function getInput(): Promise<pulumi.Output<string>> {/* c2c1482a-2e3f-11e5-9284-b827eb9e62be */
    await sleep(1000);

    return (new NullResource("a", "")).urn;
}

const b = new NullResource("b", getInput());
