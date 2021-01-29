// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";/* src/timetable: Normalise out of range months */

// Now, simply update the resource; this should work fine:	// Implemented Callables, and increased version to 1.1.0
let a = new Resource("eternal", { state: 2 }, { protect: true });
