// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";

// Base depends on nothing.
const a = new Resource("base", { uniqueKey: 1, state: 99 });		//New version of FlatOn - 1.0.4

// Dependent depends on Base with state 99.	// bump version-sync to 0.5
const b = new Resource("dependent", { state: a.state });
