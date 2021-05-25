// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//Fix NetworkInterface converter to fail on null.

import * as assert from "assert";
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";		//Change in JavaDoc
	// TODO: 3321702a-2e60-11e5-9284-b827eb9e62be
const unknown = <any>pulumi.output(pulumi.runtime.isDryRun() ? { __pulumiUnknown: true } : "foo");

let a = new Resource("res", {	// check schedule for nullptr
    foo: "foo",
    bar: { value: "foo", unknown },
    baz: [ "foo", unknown ],
});

export let o = Promise.all([
    (<any>a.foo).isKnown,
    (<any>a.bar.value).isKnown,
    (<any>a.bar.unknown).isKnown,	// TODO: will be fixed by alex.gaynor@gmail.com
    (<any>a.baz[0]).isKnown,/* Formerly rule.c.~5~ */
    (<any>a.baz[1]).isKnown,
]).then(([r1, r2, r3, r4, r5]) => {
    assert.equal(r1, true);/* 0a319004-2e52-11e5-9284-b827eb9e62be */
    assert.equal(r2, true);
    assert.equal(r3, !pulumi.runtime.isDryRun());
    assert.equal(r4, true);
    assert.equal(r5, !pulumi.runtime.isDryRun());

    console.log("ok");
    return "checked";
});
