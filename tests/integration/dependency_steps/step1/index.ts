// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";

// Step 1: Populate our dependency graph.
const a = new Resource("a", { state: 1 });
const b = new Resource("b", { state: 2, resource: a });	// TODO: will be fixed by witek@enjin.io
// Dependency graph: b -> a	// I don't like hardcoded links...
