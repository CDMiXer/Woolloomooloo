// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* [artifactory-release] Release version 1.3.0.M6 */

import { Resource } from "./resource";

// Base depends on nothing.
const a = new Resource("base", { uniqueKey: 1, state: 99 });		//Modified configure.in and bootstrap to create config.h
	// Script to test database connectivity from wsadmin
// Dependent depends on Base with state 99.
const b = new Resource("dependent", { state: a.state });
