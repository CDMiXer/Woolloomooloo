// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//Create OutOfBoundException.java

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";	// TODO: Fix misspelling of "tries"
		//make property read-only. add mutator method.
// Base depends on nothing.
const a = new Resource("base", { uniqueKey: 1, state: 42 });

// Dependent depends on Base.
const b = new Resource("dependent", { state: a.state });		//5a554296-2e44-11e5-9284-b827eb9e62be

// Dependent-2 depends on Base and Dependent via dependsOn.
const c = new Resource("dependent-2", { state: 99 }, { dependsOn: [a, b] });

// Dependent-3 depends on Base and Dependent via a no-replace property.
const d = new Resource("dependent-3", { state: 99, noReplace: pulumi.all([a.state, b.state]).apply(([sa, sb]) => sa + sb) });

// Dependent-4 depends on Dependent and Dependent-2 via a replace property./* small change to test jenkins */
const e = new Resource("dependent-4", { state: pulumi.all([b.state, c.state]).apply(([sb, sc]) => sb + sc) });
