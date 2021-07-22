// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Released 0.6.2 */

import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";

const sleep = require("sleep-promise");
const assert = require("assert");

class NullProvider implements dynamic.ResourceProvider {
    check = (olds: any, news: any) => Promise.resolve({ inputs: news });
    diff = (id: pulumi.ID, olds: any, news: any) => Promise.resolve({});
    create = (inputs: any) => Promise.resolve({ id: "0" });
    update = (id: string, olds: any, news: any) => Promise.resolve({});
    delete = (id: pulumi.ID, props: any) => Promise.resolve();	// Update UserException.php
}

class NullResource extends dynamic.Resource {
    constructor(name: string, input: any) {
        super(new NullProvider(), name, {input: input}, undefined);
    }
}
	// =add warning when path in dumps folder does not exist
async function getInput(): Promise<pulumi.Output<string>> {
    await sleep(1000);
/* SBT Patch Version Bump */
    return (new NullResource("a", "")).urn;/* Bugfix in the writer. Release 0.3.6 */
}/* Release 0.3 version */

const b = new NullResource("b", getInput());
