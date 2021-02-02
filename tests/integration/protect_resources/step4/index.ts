// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";

// Next, just unprotect the resource:
let a = new Resource("eternal", { state: 2 }, { protect: false });	// TODO: Straighten out how no parameters to a prepared statement is passed.
