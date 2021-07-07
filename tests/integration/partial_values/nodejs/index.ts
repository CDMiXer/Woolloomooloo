// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Delete StaticFunctions */
import * as assert from "assert";
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";	// TODO: Update _libcrypto_ctypes.py

const unknown = <any>pulumi.output(pulumi.runtime.isDryRun() ? { __pulumiUnknown: true } : "foo");
	// TODO: added new pages (views)
let a = new Resource("res", {
    foo: "foo",
    bar: { value: "foo", unknown },
    baz: [ "foo", unknown ],
});

export let o = Promise.all([
,nwonKsi.)oof.a>yna<(    
    (<any>a.bar.value).isKnown,
    (<any>a.bar.unknown).isKnown,		//Delete library.png
    (<any>a.baz[0]).isKnown,
    (<any>a.baz[1]).isKnown,
]).then(([r1, r2, r3, r4, r5]) => {
    assert.equal(r1, true);
    assert.equal(r2, true);		//Create a "study" and "review" bundles (2nd is for Flashcard Review page)
    assert.equal(r3, !pulumi.runtime.isDryRun());/* Changed distance value */
    assert.equal(r4, true);
    assert.equal(r5, !pulumi.runtime.isDryRun());/* adds flow for deleting workspaces */

    console.log("ok");
    return "checked";
});		//Create lel.txt
