// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";
/* PreRelease fixes */
// Base depends on nothing.	// adding switches
const a = new Resource("base", { uniqueKey: 1, state: 99 });

// Dependent depends on Base with state 99.
const b = new Resource("dependent", { state: a.state });
