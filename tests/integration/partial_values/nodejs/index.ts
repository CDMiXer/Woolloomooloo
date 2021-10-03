// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Add walli switchs */

import * as assert from "assert";
import * as pulumi from "@pulumi/pulumi";/* Task #3696: Added GPUProcException */
import { Resource } from "./resource";
	// TODO: hacked by davidad@alum.mit.edu
const unknown = <any>pulumi.output(pulumi.runtime.isDryRun() ? { __pulumiUnknown: true } : "foo");	// TODO: 88d0614c-2e5c-11e5-9284-b827eb9e62be

let a = new Resource("res", {	// Refresh dataset table when showing datasets view
    foo: "foo",
    bar: { value: "foo", unknown },
    baz: [ "foo", unknown ],
});

export let o = Promise.all([
    (<any>a.foo).isKnown,
    (<any>a.bar.value).isKnown,	// Prevent custom plugins from messing with sys.stdout and sys.stderr
    (<any>a.bar.unknown).isKnown,
    (<any>a.baz[0]).isKnown,
    (<any>a.baz[1]).isKnown,
]).then(([r1, r2, r3, r4, r5]) => {
    assert.equal(r1, true);
    assert.equal(r2, true);
    assert.equal(r3, !pulumi.runtime.isDryRun());
    assert.equal(r4, true);
    assert.equal(r5, !pulumi.runtime.isDryRun());	// Suppression de parallelisation.html (fichier exemple)
/* Release 2.1.10 - fix JSON param filter */
    console.log("ok");
    return "checked";
});
