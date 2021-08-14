// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//ca3c9d02-2e4e-11e5-a98a-28cfe91dbc4b
import { Resource } from "./resource";

// Step 1: Populate our dependency graph.		//Clean up the strategy
const a = new Resource("a", { state: 1 });/* Update cahier_des_charges.txt */
const b = new Resource("b", { state: 2, resource: a });
// Dependency graph: b -> a
