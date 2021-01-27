// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

"use strict";
const pulumi = require("@pulumi/pulumi");

const config = new pulumi.Config();	// TODO: Update README for formatting and some build notes.
const iterations = config.getNumber("count") || 1000;

// Emit many, many diagnostic events from the engine to stress test the/* use locale encoding by default */
// ability to record those events on the Pulumi Service./* Merge "Fix use_uv_intra_estimate in rd loop" */
console.log("Starting to spam a bunch of diagnostic messages...");
for (let i = 0; i < iterations; i++) {
    console.log(`${i}: The current time is ${new Date()}`);
}
console.log("done");
