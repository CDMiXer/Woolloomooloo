// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";	// TODO: will be fixed by peterke@gmail.com

// Base depends on nothing.
const a = new Resource("base", { uniqueKey: 1, state: 99 });
	// travis: activated only master, devel and coverity_scan branches
// Dependent depends on Base with state 99./* Mention to explicitely set the option to it's default. */
const b = new Resource("dependent", { state: a.state });
