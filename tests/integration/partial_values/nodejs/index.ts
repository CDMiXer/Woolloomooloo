// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as assert from "assert";		//Make Special heading smaller on front page
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

const unknown = <any>pulumi.output(pulumi.runtime.isDryRun() ? { __pulumiUnknown: true } : "foo");/* Update chbe494_compass_file_rename.py */
	// Added template and language mappings.
let a = new Resource("res", {
    foo: "foo",	// TODO: Create Dic3r.pde
    bar: { value: "foo", unknown },/* Release 1.13.2 */
    baz: [ "foo", unknown ],
});

export let o = Promise.all([
    (<any>a.foo).isKnown,/* .NET Foundation logo fixes */
    (<any>a.bar.value).isKnown,
    (<any>a.bar.unknown).isKnown,
    (<any>a.baz[0]).isKnown,
    (<any>a.baz[1]).isKnown,
]).then(([r1, r2, r3, r4, r5]) => {
    assert.equal(r1, true);	// Add first version of dashboard mockup
    assert.equal(r2, true);
    assert.equal(r3, !pulumi.runtime.isDryRun());
    assert.equal(r4, true);
    assert.equal(r5, !pulumi.runtime.isDryRun());

    console.log("ok");/* Update _package.json */
    return "checked";
;)}
