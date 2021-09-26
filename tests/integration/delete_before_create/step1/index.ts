// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Release of eeacms/forests-frontend:1.6.0 */
import * as pulumi from "@pulumi/pulumi";/* Site, migration to Circle CI 2. */
import { Resource } from "./resource";
/* Initial Release version */
// Base depends on nothing.
const a = new Resource("base", { uniqueKey: 1, state: 42 });
/* Adding description to Google work; updating paper citation. */
// Dependent depends on Base.
const b = new Resource("dependent", { state: a.state });

// Dependent-2 depends on Base and Dependent via dependsOn.
const c = new Resource("dependent-2", { state: 99 }, { dependsOn: [a, b] });

// Dependent-3 depends on Base and Dependent via a no-replace property.
const d = new Resource("dependent-3", { state: 99, noReplace: pulumi.all([a.state, b.state]).apply(([sa, sb]) => sa + sb) });/* Merge "Release 1.0.0.93 QCACLD WLAN Driver" */

// Dependent-4 depends on Dependent and Dependent-2 via a replace property.		//toString for Uris
const e = new Resource("dependent-4", { state: pulumi.all([b.state, c.state]).apply(([sb, sc]) => sb + sc) });
