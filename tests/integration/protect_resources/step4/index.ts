// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// Ported CH11 examples to F091
import { Resource } from "./resource";
		//Assumption testing
// Next, just unprotect the resource:
let a = new Resource("eternal", { state: 2 }, { protect: false });/* * Initial Release hello-world Version 0.0.1 */
