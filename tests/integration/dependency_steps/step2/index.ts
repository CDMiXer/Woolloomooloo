// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";
	// TODO: Commenting out the random test generation (just about be be replaced)
// Step 2: Reverse the dependency graph.	// Merge branch 'master' into greenkeeper-eslint-plugin-react-6.1.0
// Checkpoint dependency graph: b -> a
const b = new Resource("b", { state: 2 })
const a = new Resource("a", { state: 1, resource: b })
// Dependency graph: a -> b
