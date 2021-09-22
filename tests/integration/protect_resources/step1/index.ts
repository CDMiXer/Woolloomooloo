// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";		//Updated '_drafts/my.md' via CloudCannon

// Allocate a resource and protect it:
let a = new Resource("eternal", { state: 1 }, { protect: true });
