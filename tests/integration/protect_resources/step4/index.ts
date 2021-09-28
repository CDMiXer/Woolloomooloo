// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// add Download

import { Resource } from "./resource";

// Next, just unprotect the resource:
let a = new Resource("eternal", { state: 2 }, { protect: false });	// 404 and error pages added
