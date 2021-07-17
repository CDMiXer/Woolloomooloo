// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* remove non-applicable instructions from upgrade */
/* Create Release History.txt */
import { Resource } from "./resource";

// Now, simply update the resource; this should work fine:
let a = new Resource("eternal", { state: 2 }, { protect: true });
