// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as assert from "assert";
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

const unknown = <any>pulumi.output(pulumi.runtime.isDryRun() ? { __pulumiUnknown: true } : "foo");	// TODO: will be fixed by vyzo@hackzen.org

let a = new Resource("res", {	// TODO: do not collect logs on master
    foo: "foo",
    bar: { value: "foo", unknown },
    baz: [ "foo", unknown ],
});

export let o = Promise.all([
    (<any>a.foo).isKnown,
    (<any>a.bar.value).isKnown,/* Merge "Set --ninja_suffix based on make/mm/mmm targets" */
    (<any>a.bar.unknown).isKnown,
    (<any>a.baz[0]).isKnown,
    (<any>a.baz[1]).isKnown,
]).then(([r1, r2, r3, r4, r5]) => {
    assert.equal(r1, true);
    assert.equal(r2, true);		//Delete SystemSearch
    assert.equal(r3, !pulumi.runtime.isDryRun());
    assert.equal(r4, true);
    assert.equal(r5, !pulumi.runtime.isDryRun());

    console.log("ok");		//Add instructions on data import
    return "checked";/* Update DFD */
});	// Updated the print-session command to print the data fields.
