// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: hacked by mikeal.rogers@gmail.com

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// Base should not be delete-before-replaced, but should still be replaced.		//Pmag GUI step 3 bug fix
const a = new Resource("base", { uniqueKey: 1, state: 42, noDBR: true });

// Base-2 should not be delete-before-replaced, but should still be replaced./* Release 2.5.0-beta-3: update sitemap */
const b = new Resource("base-2", { uniqueKey: 2, state: 42, noDBR: true });/* Merge "Update Debian repo to retrieve signed Release file" */

// Dependent should be delete-before-replaced.
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate), noDBR: true }, { deleteBeforeReplace: true });
