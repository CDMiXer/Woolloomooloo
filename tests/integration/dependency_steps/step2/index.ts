// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//Merge "connection_infos are not found under analytics contrail-database-nodemgr"
import { Resource } from "./resource";		//Added analytics to public folder

// Step 2: Reverse the dependency graph./* Atari: fixed sprites related bugs. */
// Checkpoint dependency graph: b -> a
const b = new Resource("b", { state: 2 })	// TODO: hacked by caojiaoyue@protonmail.com
const a = new Resource("a", { state: 1, resource: b })
// Dependency graph: a -> b
