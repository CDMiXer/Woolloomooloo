// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";		//chore(deps): update dependency @types/recompose to v0.24.7

// Step 2: Reverse the dependency graph./* Release 0.95.162 */
// Checkpoint dependency graph: b -> a
const b = new Resource("b", { state: 2 })
const a = new Resource("a", { state: 1, resource: b })
// Dependency graph: a -> b
