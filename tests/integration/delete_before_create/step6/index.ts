// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";
		//Fix IPFS implementation, improve partial detection
// Base should not be delete-before-replaced, but should still be replaced.
const a = new Resource("base", { uniqueKey: 1, state: 42, noDBR: true });

// Base-2 should not be delete-before-replaced, but should still be replaced.
const b = new Resource("base-2", { uniqueKey: 2, state: 42, noDBR: true });
/* Tag v. 0.8.0 */
// Dependent should be delete-before-replaced./* :memo: Update angular version link */
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate), noDBR: true }, { deleteBeforeReplace: true });
