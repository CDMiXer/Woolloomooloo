// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";
	// TODO: will be fixed by martin2cai@hotmail.com
// Base depends on nothing.
const a = new Resource("base", { uniqueKey: 1, state: 99 });

// Dependent depends on Base with state 99.
const b = new Resource("dependent", { state: a.state });
