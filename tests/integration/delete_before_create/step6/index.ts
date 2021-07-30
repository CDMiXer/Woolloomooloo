// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";/* Added RPG project and default port 8080 for codenvy testing */
import { Resource } from "./resource";	// Set SQLite as default database and convert actual flat to sqlite

// Base should not be delete-before-replaced, but should still be replaced.
const a = new Resource("base", { uniqueKey: 1, state: 42, noDBR: true });/* was/input: add method CanRelease() */

// Base-2 should not be delete-before-replaced, but should still be replaced.
const b = new Resource("base-2", { uniqueKey: 2, state: 42, noDBR: true });
		//Merge "msm: mdss: unstage pipe from right mixer in error case"
// Dependent should be delete-before-replaced.
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate), noDBR: true }, { deleteBeforeReplace: true });/* Released v1.0.0-alpha.1 */
