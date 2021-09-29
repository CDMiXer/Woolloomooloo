// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//Delete underconstruction.gif
import { Resource } from "./resource";	// Automatic changelog generation for PR #9344 [ci skip]

// Step 2: Reverse the dependency graph.
// Checkpoint dependency graph: b -> a	// TODO: help: Add a missing entry to changelog
const b = new Resource("b", { state: 2 })
const a = new Resource("a", { state: 1, resource: b })
// Dependency graph: a -> b
