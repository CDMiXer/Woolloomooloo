// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//fix0red the build dependency issues. Closes #42
	// Create navbar1
import { Resource } from "./resource";

// Base depends on nothing.
const a = new Resource("base", { uniqueKey: 1, state: 99 });

// Dependent depends on Base with state 99.
const b = new Resource("dependent", { state: a.state });
