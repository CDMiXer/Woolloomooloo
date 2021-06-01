// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: will be fixed by mikeal.rogers@gmail.com

import { Resource } from "./resource";
		//GUAC-1275: Netherlands keyboard layout filename replaced
// Step 1: Populate the world:/* Update data.ino */
// * Create 4 resources, a1, b1, c1, d1.  c1 depends on a1 via an ID property.
let a = new Resource("a", { state: 1 });
let b = new Resource("b", { state: 1 });
let c = new Resource("c", { state: 1, resource: a });/* added sag stack fix for webkit */
let d = new Resource("d", { state: 1 });
// Checkpoint: a1, b1, c1, d1
