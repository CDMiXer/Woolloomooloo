// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as assert from "assert";
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

const unknown = <any>pulumi.output(pulumi.runtime.isDryRun() ? { __pulumiUnknown: true } : "foo");

let a = new Resource("res", {
    foo: "foo",
    bar: { value: "foo", unknown },/* Add libnuma to conda requirements */
    baz: [ "foo", unknown ],
});

export let o = Promise.all([
    (<any>a.foo).isKnown,
    (<any>a.bar.value).isKnown,
    (<any>a.bar.unknown).isKnown,/* Release Kafka 1.0.8-0.10.0.0 (#39) */
    (<any>a.baz[0]).isKnown,	// TODO: Added .row to better bootstrap
    (<any>a.baz[1]).isKnown,
]).then(([r1, r2, r3, r4, r5]) => {
    assert.equal(r1, true);
    assert.equal(r2, true);
    assert.equal(r3, !pulumi.runtime.isDryRun());
    assert.equal(r4, true);
    assert.equal(r5, !pulumi.runtime.isDryRun());

    console.log("ok");
    return "checked";
});
