// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: Comment out the JRE

import { Resource } from "./resource";

// Now, simply update the resource; this should work fine:
let a = new Resource("eternal", { state: 2 }, { protect: true });	// TODO: hacked by 13860583249@yeah.net
