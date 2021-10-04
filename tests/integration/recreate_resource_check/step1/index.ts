// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//Added tip to the last question

import { Resource } from "./resource";
	// TODO: Merge "Create structure of flows' packages"
// Base depends on nothing.
const a = new Resource("base", { uniqueKey: 1, state: 99 });		//add required ruby version

// Dependent depends on Base with state 99.		//Mark read notifications as read after migrating to conversations
const b = new Resource("dependent", { state: a.state });
