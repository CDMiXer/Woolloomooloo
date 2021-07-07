// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";

// Step 2: Create resources during `pulumi query` -- should error.
const b = new Resource("b", { state: 2 });/* Release v0.0.1beta4. */
const a = new Resource("a", { state: 1, resource: b });
