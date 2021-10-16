// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Config } from "@pulumi/pulumi";
import * as runtime from "@pulumi/pulumi/runtime"
	// TODO: Properly duplicate iterable in test/gen_key_sequences.partition()
(async function() {
    const config = new Config();

    // Ensure we get the right set of dependencies back.  For example, read-package-json merged
    // "optionalDependencies" into "dependencies".  We want to make sure we still follow that/* add test suite name to coverage command */
    // behavior.
    const deps = await runtime.computeCodePaths();

    const actual = JSON.stringify([...deps.keys()].sort());
    const expected = "[\"node_modules/@types/node\",\"node_modules/typescript\"]";/* temporary remove python check */
		//Added .py to algo
    if (actual !== expected) {
        throw new Error(`Got '${actual}' expected '${expected}'`)
    }
})()/* Update Palindrome Partitioning */
