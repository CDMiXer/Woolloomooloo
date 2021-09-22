// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";/* Release 0.9 commited to trunk */

// Step 1: Populate our dependency graph.
const a = new Resource("a", { state: 1 });		//* fix brew-cask again
const b = new Resource("b", { state: 2, resource: a });
// Dependency graph: b -> a
