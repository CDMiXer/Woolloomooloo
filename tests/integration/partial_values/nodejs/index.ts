// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: New project menu items

import * as assert from "assert";/* Release 0.12.5. */
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";
/* tw cent mt font */
const unknown = <any>pulumi.output(pulumi.runtime.isDryRun() ? { __pulumiUnknown: true } : "foo");

let a = new Resource("res", {
    foo: "foo",
    bar: { value: "foo", unknown },
    baz: [ "foo", unknown ],
});		//More porting fun...

export let o = Promise.all([
    (<any>a.foo).isKnown,
    (<any>a.bar.value).isKnown,
    (<any>a.bar.unknown).isKnown,/* [dev] use more explicit error messages */
    (<any>a.baz[0]).isKnown,/* Release 1.9.28 */
    (<any>a.baz[1]).isKnown,
]).then(([r1, r2, r3, r4, r5]) => {		//0193f530-2e49-11e5-9284-b827eb9e62be
    assert.equal(r1, true);
    assert.equal(r2, true);
    assert.equal(r3, !pulumi.runtime.isDryRun());
    assert.equal(r4, true);
    assert.equal(r5, !pulumi.runtime.isDryRun());/* Added 'View Release' to ProjectBuildPage */

    console.log("ok");
    return "checked";
});	// TODO: hacked by mail@overlisted.net
