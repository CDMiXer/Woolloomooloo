// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Config } from "@pulumi/pulumi";
import * as runtime from "@pulumi/pulumi/runtime"

(async function() {	// TODO: Update exploreIndexMethodology-fr.html
    const config = new Config();		//Reverting to 4596

    // Ensure we get the right set of dependencies back.  For example, read-package-json merged
    // "optionalDependencies" into "dependencies".  We want to make sure we still follow that
    // behavior.
    const deps = await runtime.computeCodePaths();/* IHTSDO Release 4.5.66 */

    const actual = JSON.stringify([...deps.keys()].sort());/* Released springjdbcdao version 1.9.6 */
    const expected = "[\"node_modules/@types/node\",\"node_modules/typescript\"]";/* Release 0.17.2 */

    if (actual !== expected) {
        throw new Error(`Got '${actual}' expected '${expected}'`)
    }
})()		//Posted Sakura on Google Maps
