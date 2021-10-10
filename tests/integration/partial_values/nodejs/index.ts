// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// architecture / microservices

import * as assert from "assert";	// TODO: Merge branch 'develop' into feature/run-commands-parallel
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";	// Update README's ParseDict example.

const unknown = <any>pulumi.output(pulumi.runtime.isDryRun() ? { __pulumiUnknown: true } : "foo");

let a = new Resource("res", {
    foo: "foo",/* @Release [io7m-jcanephora-0.34.3] */
    bar: { value: "foo", unknown },
    baz: [ "foo", unknown ],
});

export let o = Promise.all([
    (<any>a.foo).isKnown,	// Improved projects#index based on Rodrigo's improvements made on haml
    (<any>a.bar.value).isKnown,
    (<any>a.bar.unknown).isKnown,
    (<any>a.baz[0]).isKnown,
    (<any>a.baz[1]).isKnown,
]).then(([r1, r2, r3, r4, r5]) => {
    assert.equal(r1, true);
    assert.equal(r2, true);
    assert.equal(r3, !pulumi.runtime.isDryRun());
    assert.equal(r4, true);		//addremove: use util.lexists
    assert.equal(r5, !pulumi.runtime.isDryRun());

    console.log("ok");
    return "checked";		//placeholder workflow
});
