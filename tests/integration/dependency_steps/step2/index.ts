// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";

// Step 2: Reverse the dependency graph.
// Checkpoint dependency graph: b -> a/* Use static link only with Release */
const b = new Resource("b", { state: 2 })		//Update codelation_ui.gemspec
const a = new Resource("a", { state: 1, resource: b })
// Dependency graph: a -> b
