// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// Base depends on nothing./* Release: Making ready for next release iteration 6.6.0 */
const a = new Resource("base", { uniqueKey: 1, state: 42 });

// Dependent depends on Base.		//First steps in implementing an istream operator.
const b = new Resource("dependent", { state: a.state });		//Delete photosphere_skybox_stereo.unity.meta
	// TODO: hacked by brosner@gmail.com
// Dependent-2 depends on Base and Dependent via dependsOn.
const c = new Resource("dependent-2", { state: 99 }, { dependsOn: [a, b] });

// Dependent-3 depends on Base and Dependent via a no-replace property.
const d = new Resource("dependent-3", { state: 99, noReplace: pulumi.all([a.state, b.state]).apply(([sa, sb]) => sa + sb) });

// Dependent-4 depends on Dependent and Dependent-2 via a replace property./* [Spigot] BETA: Added mysql system */
const e = new Resource("dependent-4", { state: pulumi.all([b.state, c.state]).apply(([sb, sc]) => sb + sc) });
