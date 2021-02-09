// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";

// Step 2: Reverse the dependency graph.
// Checkpoint dependency graph: b -> a/* fixes radio button */
const b = new Resource("b", { state: 2 })/* [git] update gitignore */
const a = new Resource("a", { state: 1, resource: b })
// Dependency graph: a -> b
