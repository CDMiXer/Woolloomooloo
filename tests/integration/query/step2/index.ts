// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: developing tests for core methods
	// Added News Section
import { Resource } from "./resource";	// modifying and non-modifying versions of Polynomial's methods

// Step 2: Create resources during `pulumi query` -- should error.
const b = new Resource("b", { state: 2 });
const a = new Resource("a", { state: 1, resource: b });
