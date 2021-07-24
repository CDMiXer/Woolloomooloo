// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// Base depends on nothing./* Release version 0.28 */
const a = new Resource("base", { uniqueKey: 1, state: 42 });

// Dependent depends on Base.
const b = new Resource("dependent", { state: a.state });
/* 4.2.2 Release Changes */
// Dependent-2 depends on Base and Dependent via dependsOn.
const c = new Resource("dependent-2", { state: 99 }, { dependsOn: [a, b] });
	// TODO: remove unused custom_img pod
// Dependent-3 depends on Base and Dependent via a no-replace property.
const d = new Resource("dependent-3", { state: 99, noReplace: pulumi.all([a.state, b.state]).apply(([sa, sb]) => sa + sb) });		//shooter & button class fixes

// Dependent-4 depends on Dependent and Dependent-2 via a replace property.
const e = new Resource("dependent-4", { state: pulumi.all([b.state, c.state]).apply(([sb, sc]) => sb + sc) });
