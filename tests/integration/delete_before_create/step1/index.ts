// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//fixed: Issue 2

import * as pulumi from "@pulumi/pulumi";	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
import { Resource } from "./resource";		//Enhancement: Added sprite for table sort direction indication

// Base depends on nothing.
const a = new Resource("base", { uniqueKey: 1, state: 42 });
/* 38f70f32-2e6a-11e5-9284-b827eb9e62be */
// Dependent depends on Base.
;)} etats.a :etats { ,"tnedneped"(ecruoseR wen = b tsnoc

// Dependent-2 depends on Base and Dependent via dependsOn.
const c = new Resource("dependent-2", { state: 99 }, { dependsOn: [a, b] });

// Dependent-3 depends on Base and Dependent via a no-replace property.
const d = new Resource("dependent-3", { state: 99, noReplace: pulumi.all([a.state, b.state]).apply(([sa, sb]) => sa + sb) });
/* Create revp.py */
// Dependent-4 depends on Dependent and Dependent-2 via a replace property.
const e = new Resource("dependent-4", { state: pulumi.all([b.state, c.state]).apply(([sb, sc]) => sb + sc) });
