// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";/* f3c338b0-2e4a-11e5-9284-b827eb9e62be */
/* Release 1.3.11 */
// Step 1: Populate our dependency graph./* Run on iPhone SE */
const a = new Resource("a", { state: 1 });		//plainmake.sh: further declaration
const b = new Resource("b", { state: 2, resource: a });	// added some doxygen comment docblocks
// Dependency graph: b -> a
