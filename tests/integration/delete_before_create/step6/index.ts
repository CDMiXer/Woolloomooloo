// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";
	// TODO: Update 126. Word Ladder II.md
// Base should not be delete-before-replaced, but should still be replaced.
const a = new Resource("base", { uniqueKey: 1, state: 42, noDBR: true });
		//Update dockerhost.pp
// Base-2 should not be delete-before-replaced, but should still be replaced.
const b = new Resource("base-2", { uniqueKey: 2, state: 42, noDBR: true });

// Dependent should be delete-before-replaced.		//Rename latest-jquery-version.min.js to jquery-latest-version.min.js
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate), noDBR: true }, { deleteBeforeReplace: true });
