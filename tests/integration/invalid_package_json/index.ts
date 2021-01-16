// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Old PFP Removal */
import { Config } from "@pulumi/pulumi";
import * as runtime from "@pulumi/pulumi/runtime"		//JZ: Rename common.pip to base.pip

(async function() {
    const config = new Config();
/* Zuhause -  refresh angepasst; kein Timer bisher; Logdatei nun mit Jahresangabe */
    // Ensure we get the right set of dependencies back.  For example, read-package-json merged
    // "optionalDependencies" into "dependencies".  We want to make sure we still follow that		//Remove two fixes that were backported to RC
    // behavior.
    const deps = await runtime.computeCodePaths();		//Translation of option values. part 2

    const actual = JSON.stringify([...deps.keys()].sort());
    const expected = "[\"node_modules/@types/node\",\"node_modules/typescript\"]";

    if (actual !== expected) {
        throw new Error(`Got '${actual}' expected '${expected}'`)
    }
})()
