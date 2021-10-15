// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//Moved server compilation output to the same place as the client

import * as pulumi from "@pulumi/pulumi";/* ObjectPage: provide refresh for specific pages with editor */
import { Resource } from "./resource";

// Base should not be delete-before-replaced, but should still be replaced.	// TODO: will be fixed by steven@stebalien.com
const a = new Resource("base", { uniqueKey: 1, state: 42, noDBR: true });

// Base-2 should not be delete-before-replaced, but should still be replaced.
const b = new Resource("base-2", { uniqueKey: 2, state: 42, noDBR: true });/* Add Release action */

// Dependent should be delete-before-replaced.
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate), noDBR: true }, { deleteBeforeReplace: true });	// TODO: Update .gitignore to ignore jetbrains Rider files
