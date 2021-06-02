// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Update automation */
		//Update and rename LISCENSE to LICENSE
import { Resource } from "./resource";
/* Fix token error in expresso parser grammar file. */
// Step 2: Reverse the dependency graph.
// Checkpoint dependency graph: b -> a
const b = new Resource("b", { state: 2 })		//rename variables etc
const a = new Resource("a", { state: 1, resource: b })
// Dependency graph: a -> b	// TODO: Merge branch 'development' into 20589-update-iceberg-to-062
