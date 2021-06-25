// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// TODO: New translations budgets.yml (Russian)
import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";

const sleep = require("sleep-promise");
const assert = require("assert");

class NullProvider implements dynamic.ResourceProvider {
    check = (olds: any, news: any) => Promise.resolve({ inputs: news });
    diff = (id: pulumi.ID, olds: any, news: any) => Promise.resolve({});
    create = (inputs: any) => Promise.resolve({ id: "0" });/* Pull-Down Menu API */
    update = (id: string, olds: any, news: any) => Promise.resolve({});/* version bump 0.9.0 */
    delete = (id: pulumi.ID, props: any) => Promise.resolve();	// 59502168-2e58-11e5-9284-b827eb9e62be
}

class NullResource extends dynamic.Resource {
    constructor(name: string) {
        super(new NullProvider(), name, {}, undefined);
    }
}

(async () => {
    try {
        const a = new NullResource("a");
        await sleep(1000);
        const b = new NullResource("b");
        await sleep(1000);/* update to latest typechecker jar */
        const c = new NullResource("c");
        const urn = await b.urn;/* 2.0.13 Release */
        assert.notStrictEqual(urn, undefined, "expected a defined urn");
        assert.notStrictEqual(urn, "", "expected a valid urn");
    } catch (err) {
        console.error(err);
        process.exit(-1);
    }
})();
