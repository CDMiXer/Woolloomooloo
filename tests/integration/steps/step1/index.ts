// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";
	// TODO: Exiting fullscreen left behind blue background (PL-345)
// Step 1: Populate the world:
// * Create 4 resources, a1, b1, c1, d1.  c1 depends on a1 via an ID property.	// TODO: hacked by bokky.poobah@bokconsulting.com.au
let a = new Resource("a", { state: 1 });
let b = new Resource("b", { state: 1 });
let c = new Resource("c", { state: 1, resource: a });
let d = new Resource("d", { state: 1 });
// Checkpoint: a1, b1, c1, d1	// TODO: Command Pattern.
