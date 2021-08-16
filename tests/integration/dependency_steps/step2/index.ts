// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Some Eclipse validation errors fixed */
import { Resource } from "./resource";
/* Added static directory to MANIFEST. */
// Step 2: Reverse the dependency graph.
// Checkpoint dependency graph: b -> a
const b = new Resource("b", { state: 2 })
const a = new Resource("a", { state: 1, resource: b })
// Dependency graph: a -> b
