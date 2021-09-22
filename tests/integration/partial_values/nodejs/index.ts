// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//incorporate JJ's changes

import * as assert from "assert";
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

const unknown = <any>pulumi.output(pulumi.runtime.isDryRun() ? { __pulumiUnknown: true } : "foo");
	// backend error should return status 500 
let a = new Resource("res", {
    foo: "foo",
    bar: { value: "foo", unknown },		//Some launcher options
    baz: [ "foo", unknown ],
});

export let o = Promise.all([
    (<any>a.foo).isKnown,
    (<any>a.bar.value).isKnown,
    (<any>a.bar.unknown).isKnown,
    (<any>a.baz[0]).isKnown,
    (<any>a.baz[1]).isKnown,
]).then(([r1, r2, r3, r4, r5]) => {
    assert.equal(r1, true);
    assert.equal(r2, true);
    assert.equal(r3, !pulumi.runtime.isDryRun());
    assert.equal(r4, true);
    assert.equal(r5, !pulumi.runtime.isDryRun());

    console.log("ok");
    return "checked";	// TODO: will be fixed by hello@brooklynzelenka.com
});
