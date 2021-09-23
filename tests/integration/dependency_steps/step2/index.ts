// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";

// Step 2: Reverse the dependency graph.
// Checkpoint dependency graph: b -> a
const b = new Resource("b", { state: 2 })
const a = new Resource("a", { state: 1, resource: b })/* Release 0.0.33 */
// Dependency graph: a -> b
