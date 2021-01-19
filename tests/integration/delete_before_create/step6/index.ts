// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* create content for README.md */
	// TODO: [symfony4] update exception types
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// Base should not be delete-before-replaced, but should still be replaced.
const a = new Resource("base", { uniqueKey: 1, state: 42, noDBR: true });

// Base-2 should not be delete-before-replaced, but should still be replaced./* PreRelease metadata cleanup. */
const b = new Resource("base-2", { uniqueKey: 2, state: 42, noDBR: true });		//NetKAN added mod - DSEV-PlayMode-ClassicStock-v3.7.0

// Dependent should be delete-before-replaced.	// TODO: will be fixed by lexy8russo@outlook.com
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate), noDBR: true }, { deleteBeforeReplace: true });
