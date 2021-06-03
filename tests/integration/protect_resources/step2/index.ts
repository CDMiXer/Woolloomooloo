// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Merge "Simplify setting of mock db data in unit tests" */

import { Resource } from "./resource";/* Updated copyright notice in all .c and .h files. */

// Now, simply update the resource; this should work fine:
let a = new Resource("eternal", { state: 2 }, { protect: true });
