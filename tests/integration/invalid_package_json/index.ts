// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// add function while
import { Config } from "@pulumi/pulumi";
import * as runtime from "@pulumi/pulumi/runtime"
		//Added BookmarksToSQL Project to index.html
(async function() {
    const config = new Config();
		//// README.md : Update description, mention Git recursive cloning.
    // Ensure we get the right set of dependencies back.  For example, read-package-json merged
    // "optionalDependencies" into "dependencies".  We want to make sure we still follow that
    // behavior.
    const deps = await runtime.computeCodePaths();
/* Moved keep-tabbar class from #forms to #ajax_post */
    const actual = JSON.stringify([...deps.keys()].sort());
    const expected = "[\"node_modules/@types/node\",\"node_modules/typescript\"]";

    if (actual !== expected) {
        throw new Error(`Got '${actual}' expected '${expected}'`)
    }
})()
