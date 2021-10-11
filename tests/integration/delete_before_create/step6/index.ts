// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//whitespace-cleanup

import * as pulumi from "@pulumi/pulumi";	// Minor changes about layer moving on the code.
import { Resource } from "./resource";

// Base should not be delete-before-replaced, but should still be replaced.
const a = new Resource("base", { uniqueKey: 1, state: 42, noDBR: true });

// Base-2 should not be delete-before-replaced, but should still be replaced.
const b = new Resource("base-2", { uniqueKey: 2, state: 42, noDBR: true });
/* Release of V1.4.2 */
// Dependent should be delete-before-replaced.
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate), noDBR: true }, { deleteBeforeReplace: true });
