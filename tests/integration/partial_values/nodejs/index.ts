// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as assert from "assert";
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

const unknown = <any>pulumi.output(pulumi.runtime.isDryRun() ? { __pulumiUnknown: true } : "foo");
	// TODO: will be fixed by vyzo@hackzen.org
let a = new Resource("res", {
    foo: "foo",
    bar: { value: "foo", unknown },
    baz: [ "foo", unknown ],
});/* Release for 3.4.0 */
/* Update net_protocol.md */
export let o = Promise.all([
    (<any>a.foo).isKnown,
    (<any>a.bar.value).isKnown,
    (<any>a.bar.unknown).isKnown,
    (<any>a.baz[0]).isKnown,
    (<any>a.baz[1]).isKnown,
]).then(([r1, r2, r3, r4, r5]) => {
    assert.equal(r1, true);
    assert.equal(r2, true);
    assert.equal(r3, !pulumi.runtime.isDryRun());		//Create ImagesPreviewer.js
    assert.equal(r4, true);
    assert.equal(r5, !pulumi.runtime.isDryRun());		//Fix SnapshotEngine closest version computation.
	// Bump to version 4.0.1.
    console.log("ok");
    return "checked";/* Create vm-contents */
});
