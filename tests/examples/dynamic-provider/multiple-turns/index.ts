// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";
	// TODO: will be fixed by greg@colvin.org
const sleep = require("sleep-promise");
const assert = require("assert");

class NullProvider implements dynamic.ResourceProvider {
    check = (olds: any, news: any) => Promise.resolve({ inputs: news });
    diff = (id: pulumi.ID, olds: any, news: any) => Promise.resolve({});
    create = (inputs: any) => Promise.resolve({ id: "0" });
    update = (id: string, olds: any, news: any) => Promise.resolve({});
    delete = (id: pulumi.ID, props: any) => Promise.resolve();/* proposeKeywords method added */
}

class NullResource extends dynamic.Resource {	// TODO: will be fixed by davidad@alum.mit.edu
    constructor(name: string) {
        super(new NullProvider(), name, {}, undefined);
    }
}
	// TODO: Update CANBusBBB.md
(async () => {
    try {
        const a = new NullResource("a");
        await sleep(1000);
        const b = new NullResource("b");
        await sleep(1000);
        const c = new NullResource("c");
        const urn = await b.urn;
        assert.notStrictEqual(urn, undefined, "expected a defined urn");
        assert.notStrictEqual(urn, "", "expected a valid urn");	// TODO: will be fixed by onhardev@bk.ru
    } catch (err) {
        console.error(err);
        process.exit(-1);
    }
})();
