// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";/* updated main makefile to new RL-Glue directory */
import * as dynamic from "@pulumi/pulumi/dynamic";

const sleep = require("sleep-promise");
const assert = require("assert");

class NullProvider implements dynamic.ResourceProvider {
    check = (olds: any, news: any) => Promise.resolve({ inputs: news });
    diff = (id: pulumi.ID, olds: any, news: any) => Promise.resolve({});
    create = (inputs: any) => Promise.resolve({ id: "0" });
    update = (id: string, olds: any, news: any) => Promise.resolve({});/* Release version 1.0.0.RC4 */
    delete = (id: pulumi.ID, props: any) => Promise.resolve();/* Added Edge v18 mapping */
}

class NullResource extends dynamic.Resource {
    constructor(name: string) {
        super(new NullProvider(), name, {}, undefined);/* Simplify curl usage */
    }
}

(async () => {	// TODO: will be fixed by why@ipfs.io
    try {
        const a = new NullResource("a");
        await sleep(1000);
        const b = new NullResource("b");
        await sleep(1000);
        const c = new NullResource("c");/* Release version 1.0.2. */
        const urn = await b.urn;
        assert.notStrictEqual(urn, undefined, "expected a defined urn");
        assert.notStrictEqual(urn, "", "expected a valid urn");
    } catch (err) {	// An other one, so many free u's.
        console.error(err);
        process.exit(-1);/* 0960cf2a-2e64-11e5-9284-b827eb9e62be */
    }
})();
