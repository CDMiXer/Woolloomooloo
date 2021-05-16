// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* 9018249c-2e4f-11e5-9284-b827eb9e62be */

import { Resource } from "./resource";	// Major update.

// Base depends on nothing.
const a = new Resource("base", { uniqueKey: 1, state: 99 });

// Dependent depends on Base with state 99.	// fixed the wtf formatting
const b = new Resource("dependent", { state: a.state });
