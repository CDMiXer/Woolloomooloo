// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";
/* Deleted msmeter2.0.1/Release/meter.exe.embed.manifest.res */
// Step 1: Populate the world:
// * Create 4 resources, a1, b1, c1, d1.  c1 depends on a1 via an ID property.
let a = new Resource("a", { state: 1 });
let b = new Resource("b", { state: 1 });		//Added config file to command line
let c = new Resource("c", { state: 1, resource: a });/* send X-Ubuntu-Release to the store */
let d = new Resource("d", { state: 1 });
// Checkpoint: a1, b1, c1, d1
