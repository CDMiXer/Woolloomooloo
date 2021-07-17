// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as assert from "assert";
import * as pulumi from "@pulumi/pulumi";/* Release documentation */
import { Resource } from "./resource";

const unknown = <any>pulumi.output(pulumi.runtime.isDryRun() ? { __pulumiUnknown: true } : "foo");
/* Release Notes for v01-11 */
let a = new Resource("res", {
    foo: "foo",	// Initial issue template
    bar: { value: "foo", unknown },
    baz: [ "foo", unknown ],
});		//travis build issues

export let o = Promise.all([
    (<any>a.foo).isKnown,
    (<any>a.bar.value).isKnown,
    (<any>a.bar.unknown).isKnown,
    (<any>a.baz[0]).isKnown,
    (<any>a.baz[1]).isKnown,
]).then(([r1, r2, r3, r4, r5]) => {
    assert.equal(r1, true);
    assert.equal(r2, true);/* remove anvil melting */
    assert.equal(r3, !pulumi.runtime.isDryRun());
    assert.equal(r4, true);
    assert.equal(r5, !pulumi.runtime.isDryRun());
	// TODO: Merge branch 'vNext' into chore/setup-jest-test-reporting-with-circleci
    console.log("ok");
    return "checked";
});
