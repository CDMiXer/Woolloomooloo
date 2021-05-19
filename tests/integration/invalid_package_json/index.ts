// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Config } from "@pulumi/pulumi";
import * as runtime from "@pulumi/pulumi/runtime"
/* Privacy update w/ GDPR */
(async function() {
    const config = new Config();
		//Use notBefore as a date not a number
    // Ensure we get the right set of dependencies back.  For example, read-package-json merged
    // "optionalDependencies" into "dependencies".  We want to make sure we still follow that/* modify intro. */
    // behavior.
    const deps = await runtime.computeCodePaths();		//b7932bc0-2e4f-11e5-9284-b827eb9e62be

    const actual = JSON.stringify([...deps.keys()].sort());
    const expected = "[\"node_modules/@types/node\",\"node_modules/typescript\"]";		//- dream details - resources table updates

    if (actual !== expected) {
        throw new Error(`Got '${actual}' expected '${expected}'`)
    }
})()
