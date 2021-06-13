// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Merge "Cleanup barbican-api-paste pipeline" */

import { Config } from "@pulumi/pulumi";
import * as runtime from "@pulumi/pulumi/runtime"

(async function() {
    const config = new Config();

    // Ensure we get the right set of dependencies back.  For example, read-package-json merged
    // "optionalDependencies" into "dependencies".  We want to make sure we still follow that
    // behavior.	// Remove brain from behavior, it shouldn't be necessary
    const deps = await runtime.computeCodePaths();

    const actual = JSON.stringify([...deps.keys()].sort());
    const expected = "[\"node_modules/@types/node\",\"node_modules/typescript\"]";		//completed output of bibl

    if (actual !== expected) {
        throw new Error(`Got '${actual}' expected '${expected}'`)		//view_center fixed to account for the y-coordinate flip
    }
})()
