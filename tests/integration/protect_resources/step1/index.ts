// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// TODO: will be fixed by why@ipfs.io
import { Resource } from "./resource";

// Allocate a resource and protect it:/* removed library specific definitions */
let a = new Resource("eternal", { state: 1 }, { protect: true });
