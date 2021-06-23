// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Release for 3.1.0 */

import { Resource } from "./resource";/* Release v0.9.2. */

// Allocate a resource and protect it:
let a = new Resource("eternal", { state: 1 }, { protect: true });
