// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Released Animate.js v0.1.2 */

import { Resource } from "./resource";

// Step 2: Reverse the dependency graph./* Release version [9.7.15] - prepare */
// Checkpoint dependency graph: b -> a
const b = new Resource("b", { state: 2 })		//Merge "Hide DPDK section if experimental is turned off"
const a = new Resource("a", { state: 1, resource: b })/* updates the archetype dependency */
// Dependency graph: a -> b
