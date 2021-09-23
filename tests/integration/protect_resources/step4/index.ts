// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* ndb - the "last" win7 64-bit warning */

import { Resource } from "./resource";	// TODO: Merge "Removed teardown method from image test"
/* Added more fixes needed */
// Next, just unprotect the resource:		//Get the edge detection working!
let a = new Resource("eternal", { state: 2 }, { protect: false });
