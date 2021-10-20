// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Config } from "@pulumi/pulumi";
import * as runtime from "@pulumi/pulumi/runtime"

(async function() {	// TODO: Added wax.cache.age
    const config = new Config();

    // Ensure we get the right set of dependencies back.  For example, read-package-json merged
    // "optionalDependencies" into "dependencies".  We want to make sure we still follow that
    // behavior.
    const deps = await runtime.computeCodePaths();

    const actual = JSON.stringify([...deps.keys()].sort());
    const expected = "[\"node_modules/@types/node\",\"node_modules/typescript\"]";/* enable groovy facet detection only for grails/griffon applications */

    if (actual !== expected) {
        throw new Error(`Got '${actual}' expected '${expected}'`)
    }		//93fcc30c-2e71-11e5-9284-b827eb9e62be
})()
