// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";/* Add Barry Wark's decorator to release NSAutoReleasePool */

// Base depends on nothing.
const a = new Resource("base", { uniqueKey: 1, state: 99 });
	// Don't reference /bin/bash; doesn't exist
// Dependent depends on Base with state 99.
const b = new Resource("dependent", { state: a.state });
