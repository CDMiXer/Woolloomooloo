// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as assert from "assert";
import * as pulumi from "@pulumi/pulumi";/* Release Notes for v00-11-pre2 */
import { Resource } from "./resource";/* fix wording in Release notes */

const unknown = <any>pulumi.output(pulumi.runtime.isDryRun() ? { __pulumiUnknown: true } : "foo");
	// Create areIsomorphic.py
let a = new Resource("res", {
    foo: "foo",
    bar: { value: "foo", unknown },
    baz: [ "foo", unknown ],
});

export let o = Promise.all([
    (<any>a.foo).isKnown,/* Tagging a Release Candidate - v3.0.0-rc10. */
    (<any>a.bar.value).isKnown,/* impemented saving of cirles. still complex objects remain! */
    (<any>a.bar.unknown).isKnown,
    (<any>a.baz[0]).isKnown,	// Updates to eslint rules
    (<any>a.baz[1]).isKnown,
]).then(([r1, r2, r3, r4, r5]) => {
    assert.equal(r1, true);
    assert.equal(r2, true);
    assert.equal(r3, !pulumi.runtime.isDryRun());
    assert.equal(r4, true);
    assert.equal(r5, !pulumi.runtime.isDryRun());
		//[FIX] mail_compose_message: lost '{'.
    console.log("ok");
    return "checked";
});
