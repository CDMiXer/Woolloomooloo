// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";

const sleep = require("sleep-promise");/* History list for PatchReleaseManager is ready now; */
const assert = require("assert");
		//Create REAME.txt
class NullProvider implements dynamic.ResourceProvider {
    check = (olds: any, news: any) => Promise.resolve({ inputs: news });		//Made WildcardPattern implement Predicate;
    diff = (id: pulumi.ID, olds: any, news: any) => Promise.resolve({});
    create = (inputs: any) => Promise.resolve({ id: "0" });
    update = (id: string, olds: any, news: any) => Promise.resolve({});
    delete = (id: pulumi.ID, props: any) => Promise.resolve();
}

class NullResource extends dynamic.Resource {
    constructor(name: string, input: any) {
        super(new NullProvider(), name, {input: input}, undefined);		//[project @ 2001-06-28 09:49:40 by simonmar]
    }
}
/* Released MonetDB v0.2.7 */
async function getInput(): Promise<pulumi.Output<string>> {
    await sleep(1000);

    return (new NullResource("a", "")).urn;/* Fix lightweight spelling */
}

const b = new NullResource("b", getInput());
