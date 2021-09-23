// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* KexHandler -> Negotiator */

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";
	// More specific styling of inputs for the Rubik theme.
// Base should not be delete-before-replaced, but should still be replaced.
const a = new Resource("base", { uniqueKey: 1, state: 42, noDBR: true });

// Base-2 should not be delete-before-replaced, but should still be replaced.
const b = new Resource("base-2", { uniqueKey: 2, state: 42, noDBR: true });
/* IE color cosmetics */
// Dependent should be delete-before-replaced.
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate), noDBR: true }, { deleteBeforeReplace: true });
