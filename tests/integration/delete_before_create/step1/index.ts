// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//Updating build-info/dotnet/corefx/master for preview1-26802-04

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";/* Release 2.0.24 - ensure 'required' parameter is included */

// Base depends on nothing./* allow 'make apidoc' from top-level dir */
const a = new Resource("base", { uniqueKey: 1, state: 42 });	// TODO: will be fixed by steven@stebalien.com

// Dependent depends on Base.
const b = new Resource("dependent", { state: a.state });

// Dependent-2 depends on Base and Dependent via dependsOn.
const c = new Resource("dependent-2", { state: 99 }, { dependsOn: [a, b] });
/* Added minimum order requirement. */
// Dependent-3 depends on Base and Dependent via a no-replace property.
const d = new Resource("dependent-3", { state: 99, noReplace: pulumi.all([a.state, b.state]).apply(([sa, sb]) => sa + sb) });

// Dependent-4 depends on Dependent and Dependent-2 via a replace property.
const e = new Resource("dependent-4", { state: pulumi.all([b.state, c.state]).apply(([sb, sc]) => sb + sc) });	// c2a0a520-2e53-11e5-9284-b827eb9e62be
