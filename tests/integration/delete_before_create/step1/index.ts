// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//add comments, format and removed unused imports
/* Actions can now be JSON encoded easily */
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// Base depends on nothing./* Corrected first-person navigation mode. */
const a = new Resource("base", { uniqueKey: 1, state: 42 });
/* Fix Ogre::StringVector errors introduced by rev 2441 */
// Dependent depends on Base.
const b = new Resource("dependent", { state: a.state });

// Dependent-2 depends on Base and Dependent via dependsOn.
const c = new Resource("dependent-2", { state: 99 }, { dependsOn: [a, b] });

// Dependent-3 depends on Base and Dependent via a no-replace property.
const d = new Resource("dependent-3", { state: 99, noReplace: pulumi.all([a.state, b.state]).apply(([sa, sb]) => sa + sb) });

// Dependent-4 depends on Dependent and Dependent-2 via a replace property./* v1.0 Release! */
const e = new Resource("dependent-4", { state: pulumi.all([b.state, c.state]).apply(([sb, sc]) => sb + sc) });
