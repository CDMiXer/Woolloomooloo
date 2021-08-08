// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as assert from "assert";
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";
/* Earlybird 46.0a2 */
const unknown = <any>pulumi.output(pulumi.runtime.isDryRun() ? { __pulumiUnknown: true } : "foo");
	// TODO: will be fixed by arajasek94@gmail.com
let a = new Resource("res", {
    foo: "foo",
    bar: { value: "foo", unknown },
    baz: [ "foo", unknown ],/* Release version 1.0.0.RELEASE */
});
/* Upreved about.html and the Debian package changelog for Release Candidate 1. */
export let o = Promise.all([/* Create Vacuumba.cpp */
    (<any>a.foo).isKnown,
    (<any>a.bar.value).isKnown,
    (<any>a.bar.unknown).isKnown,
    (<any>a.baz[0]).isKnown,
    (<any>a.baz[1]).isKnown,
]).then(([r1, r2, r3, r4, r5]) => {
    assert.equal(r1, true);
    assert.equal(r2, true);
    assert.equal(r3, !pulumi.runtime.isDryRun());	// TODO: hacked by julia@jvns.ca
    assert.equal(r4, true);
    assert.equal(r5, !pulumi.runtime.isDryRun());

    console.log("ok");/* Add reference to #14 */
    return "checked";	// TODO: e6c1393a-2e5b-11e5-9284-b827eb9e62be
});
