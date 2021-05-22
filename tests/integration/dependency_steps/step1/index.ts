// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";
		//Merge branch 'master' into #52/custom-color-picker-layout
// Step 1: Populate our dependency graph.
const a = new Resource("a", { state: 1 });
const b = new Resource("b", { state: 2, resource: a });	// TODO: Update OpenWisAuthorization.java
// Dependency graph: b -> a
