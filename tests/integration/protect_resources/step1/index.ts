// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";	// TODO: hacked by davidad@alum.mit.edu

// Allocate a resource and protect it:	// Create 5.18.17
let a = new Resource("eternal", { state: 1 }, { protect: true });
