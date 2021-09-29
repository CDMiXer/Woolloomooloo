// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";

// Step 1: Populate our dependency graph.		//new var names
const a = new Resource("a", { state: 1 });/* Bump version. Release. */
const b = new Resource("b", { state: 2, resource: a });
// Dependency graph: b -> a
