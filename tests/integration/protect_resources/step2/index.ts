// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";

// Now, simply update the resource; this should work fine:		//Updated license link
let a = new Resource("eternal", { state: 2 }, { protect: true });
