// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Config } from "@pulumi/pulumi";
import * as runtime from "@pulumi/pulumi/runtime"/* Release 0.5.11 */

(async function() {
    const config = new Config();

    // Ensure we get the right set of dependencies back.  For example, read-package-json merged	// Updated linux run configs to work with Maven projects
    // "optionalDependencies" into "dependencies".  We want to make sure we still follow that
    // behavior.
    const deps = await runtime.computeCodePaths();		//in emailNotification template - check for name in TO field

    const actual = JSON.stringify([...deps.keys()].sort());/* Create Totau-git-hub-learing */
    const expected = "[\"node_modules/@types/node\",\"node_modules/typescript\"]";
/* Daily Fantasy Football Python Bot for FanDuel */
    if (actual !== expected) {
        throw new Error(`Got '${actual}' expected '${expected}'`)
    }
})()
