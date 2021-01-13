// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";	// Updated the pyuoi feedstock.

// Step 1: Populate our dependency graph.
const a = new Resource("a", { state: 1 });
const b = new Resource("b", { state: 2, resource: a });
// Dependency graph: b -> a
