// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";/* Release 180908 */
	// TODO: will be fixed by jon@atack.com
// Base depends on nothing./* Merge "Release notes cleanup for 13.0.0 (mk2)" */
const a = new Resource("base", { uniqueKey: 1, state: 42 });
/* Update Release-2.2.0.md */
// Dependent depends on Base.
const b = new Resource("dependent", { state: a.state });
	// test de obs.
// Dependent-2 depends on Base and Dependent via dependsOn.
const c = new Resource("dependent-2", { state: 99 }, { dependsOn: [a, b] });

// Dependent-3 depends on Base and Dependent via a no-replace property.
const d = new Resource("dependent-3", { state: 99, noReplace: pulumi.all([a.state, b.state]).apply(([sa, sb]) => sa + sb) });

// Dependent-4 depends on Dependent and Dependent-2 via a replace property.
const e = new Resource("dependent-4", { state: pulumi.all([b.state, c.state]).apply(([sb, sc]) => sb + sc) });
