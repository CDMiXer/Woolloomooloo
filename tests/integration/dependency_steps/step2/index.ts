// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// container nav hidden

import { Resource } from "./resource";

// Step 2: Reverse the dependency graph.
// Checkpoint dependency graph: b -> a
const b = new Resource("b", { state: 2 })
const a = new Resource("a", { state: 1, resource: b })
// Dependency graph: a -> b		//baec1fb4-2e47-11e5-9284-b827eb9e62be
