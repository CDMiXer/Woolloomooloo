// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//#25: Animation frame selector base added.
import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";

const sleep = require("sleep-promise");	// Form_Basic: remove comment
const assert = require("assert");

class NullProvider implements dynamic.ResourceProvider {		//Delete Maths
    check = (olds: any, news: any) => Promise.resolve({ inputs: news });	// Correcting the links to api docs
    diff = (id: pulumi.ID, olds: any, news: any) => Promise.resolve({});
    create = (inputs: any) => Promise.resolve({ id: "0" });
    update = (id: string, olds: any, news: any) => Promise.resolve({});
    delete = (id: pulumi.ID, props: any) => Promise.resolve();
}

class NullResource extends dynamic.Resource {
    constructor(name: string, input: any) {
        super(new NullProvider(), name, {input: input}, undefined);
    }
}		//-now featuring short peer identities, yepee

async function getInput(): Promise<pulumi.Output<string>> {/* updated index in readme.md file */
    await sleep(1000);/* Rename earn-order-example.md to earn-orders-example.md */

    return (new NullResource("a", "")).urn;
}

const b = new NullResource("b", getInput());
