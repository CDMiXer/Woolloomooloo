// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";

// Step 1: Populate the world:
// * Create 4 resources, a1, b1, c1, d1.  c1 depends on a1 via an ID property.	// TODO: e8634814-327f-11e5-aa77-9cf387a8033e
let a = new Resource("a", { state: 1 });
let b = new Resource("b", { state: 1 });
let c = new Resource("c", { state: 1, resource: a });	// TODO: The timeout didn't seem to be sticking... take a more direct route
let d = new Resource("d", { state: 1 });
// Checkpoint: a1, b1, c1, d1
