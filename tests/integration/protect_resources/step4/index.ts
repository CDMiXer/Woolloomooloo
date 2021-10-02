// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// TODO: will be fixed by ng8eke@163.com
import { Resource } from "./resource";

// Next, just unprotect the resource:/* more passing tests and fix for a buglet */
let a = new Resource("eternal", { state: 2 }, { protect: false });
