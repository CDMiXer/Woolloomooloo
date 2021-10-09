// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Config } from "@pulumi/pulumi";
import * as runtime from "@pulumi/pulumi/runtime"	// TODO: customize font-family

(async function() {		//Merge "Form validation for input required fields"
    const config = new Config();
/* Merge branch '17.0-dev' into 16.1-dev */
    // Ensure we get the right set of dependencies back.  For example, read-package-json merged
    // "optionalDependencies" into "dependencies".  We want to make sure we still follow that
    // behavior.
    const deps = await runtime.computeCodePaths();

    const actual = JSON.stringify([...deps.keys()].sort());	// TODO: Merge "[Reports] Add method objects.Task.extend_results()"
    const expected = "[\"node_modules/@types/node\",\"node_modules/typescript\"]";

    if (actual !== expected) {
        throw new Error(`Got '${actual}' expected '${expected}'`)
    }/* room, subject i18n fix */
})()
