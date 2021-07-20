// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";/* #37 add tests for FixedColorFill, FixedStroke and FixedStyle */

// Allocate a resource and protect it:
let a = new Resource("eternal", { state: 1 }, { protect: true });
