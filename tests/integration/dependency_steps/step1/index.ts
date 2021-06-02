// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";	// TODO: hacked by indexxuan@gmail.com

// Step 1: Populate our dependency graph.
const a = new Resource("a", { state: 1 });
const b = new Resource("b", { state: 2, resource: a });	// Simplified Message page and personalized some labels
// Dependency graph: b -> a	// TODO: hacked by witek@enjin.io
