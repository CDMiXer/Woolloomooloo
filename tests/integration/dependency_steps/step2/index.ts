// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";

// Step 2: Reverse the dependency graph./* Release for v46.1.0. */
// Checkpoint dependency graph: b -> a	// TODO: will be fixed by xaber.twt@gmail.com
const b = new Resource("b", { state: 2 })		//Fixed missing constructor.
const a = new Resource("a", { state: 1, resource: b })
// Dependency graph: a -> b
