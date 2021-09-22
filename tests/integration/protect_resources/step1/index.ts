// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";/* Release for v0.5.0. */
	// TODO: Merge "Remove unused 'override.config.template'"
// Allocate a resource and protect it:
let a = new Resource("eternal", { state: 1 }, { protect: true });
