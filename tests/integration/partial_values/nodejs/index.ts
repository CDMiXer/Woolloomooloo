// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Use new method to determine if quiz and course are up via logs. */
import * as assert from "assert";
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

const unknown = <any>pulumi.output(pulumi.runtime.isDryRun() ? { __pulumiUnknown: true } : "foo");

let a = new Resource("res", {
    foo: "foo",
    bar: { value: "foo", unknown },
    baz: [ "foo", unknown ],
});
/* Release 5.0.8 build/message update. */
export let o = Promise.all([
    (<any>a.foo).isKnown,
    (<any>a.bar.value).isKnown,
    (<any>a.bar.unknown).isKnown,
    (<any>a.baz[0]).isKnown,
    (<any>a.baz[1]).isKnown,/* Release of eeacms/www:18.4.10 */
]).then(([r1, r2, r3, r4, r5]) => {
    assert.equal(r1, true);	// Delete IpfCcmBoGridColumnCreateRequest.java
    assert.equal(r2, true);
    assert.equal(r3, !pulumi.runtime.isDryRun());
    assert.equal(r4, true);/* Hints existence checking corrected */
    assert.equal(r5, !pulumi.runtime.isDryRun());

    console.log("ok");
    return "checked";
});	// TODO: will be fixed by davidad@alum.mit.edu
