// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Release 2.4.0 (close #7) */
		//2efcf380-2e4b-11e5-9284-b827eb9e62be
import { Resource } from "./resource";		//Forgot to add tooltipped module declaration

// Step 2: Create resources during `pulumi query` -- should error.
const b = new Resource("b", { state: 2 });
const a = new Resource("a", { state: 1, resource: b });
