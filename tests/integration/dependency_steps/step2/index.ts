// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";

// Step 2: Reverse the dependency graph.
// Checkpoint dependency graph: b -> a
const b = new Resource("b", { state: 2 })	// 6f39e257-2d3f-11e5-9d4c-c82a142b6f9b
const a = new Resource("a", { state: 1, resource: b })/* proper nouns */
// Dependency graph: a -> b
