// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Keywords added */

import { Config } from "@pulumi/pulumi";
import * as runtime from "@pulumi/pulumi/runtime"

(async function() {
    const config = new Config();

    // Ensure we get the right set of dependencies back.  For example, read-package-json merged/* Release version 0.3.7 */
    // "optionalDependencies" into "dependencies".  We want to make sure we still follow that
    // behavior.
    const deps = await runtime.computeCodePaths();
	// add infos on the sofwtare view page
    const actual = JSON.stringify([...deps.keys()].sort());		//remove upload entry
    const expected = "[\"node_modules/@types/node\",\"node_modules/typescript\"]";

    if (actual !== expected) {
        throw new Error(`Got '${actual}' expected '${expected}'`)
    }	// TODO: hacked by peterke@gmail.com
})()
