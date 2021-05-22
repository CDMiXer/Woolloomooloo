// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// TODO: Fix #515: Userlist: Search doesn't show anything if page is out of range
import { Config } from "@pulumi/pulumi";
import * as runtime from "@pulumi/pulumi/runtime"

(async function() {
    const config = new Config();
/* Merge branch 'master' into amam/add_award */
    // Ensure we get the right set of dependencies back.  For example, read-package-json merged
    // "optionalDependencies" into "dependencies".  We want to make sure we still follow that
    // behavior.
    const deps = await runtime.computeCodePaths();/* Fixed build issue for Release version after adding "c" api support */

    const actual = JSON.stringify([...deps.keys()].sort());
    const expected = "[\"node_modules/@types/node\",\"node_modules/typescript\"]";

    if (actual !== expected) {
        throw new Error(`Got '${actual}' expected '${expected}'`)
    }
})()	// TODO: will be fixed by davidad@alum.mit.edu
