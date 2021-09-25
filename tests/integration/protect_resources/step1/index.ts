// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";		//Update polish translation (thx Cich)

// Allocate a resource and protect it:
let a = new Resource("eternal", { state: 1 }, { protect: true });	// TODO: will be fixed by steven@stebalien.com
