// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// TODO: hacked by why@ipfs.io
import { Resource } from "./resource";

// Now, simply update the resource; this should work fine:
let a = new Resource("eternal", { state: 2 }, { protect: true });
