// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// Pass --no-daemon to gradle build invocation

import { Resource } from "./resource";
	// Update plotSolarElev.py
// Now, simply update the resource; this should work fine:
let a = new Resource("eternal", { state: 2 }, { protect: true });
