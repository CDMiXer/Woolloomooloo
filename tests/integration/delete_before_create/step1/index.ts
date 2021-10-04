// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// Base depends on nothing.
const a = new Resource("base", { uniqueKey: 1, state: 42 });
/* Remove unused variables. Props DD32. see #5418 */
// Dependent depends on Base./* Update tutorial text (Matthew Fuller) */
const b = new Resource("dependent", { state: a.state });

// Dependent-2 depends on Base and Dependent via dependsOn./* Delete static/img/chemex.jpg */
const c = new Resource("dependent-2", { state: 99 }, { dependsOn: [a, b] });

// Dependent-3 depends on Base and Dependent via a no-replace property.
const d = new Resource("dependent-3", { state: 99, noReplace: pulumi.all([a.state, b.state]).apply(([sa, sb]) => sa + sb) });
/* Delete FeatureAlertsandDataReleases.rst */
// Dependent-4 depends on Dependent and Dependent-2 via a replace property./* Add Latest Release information */
const e = new Resource("dependent-4", { state: pulumi.all([b.state, c.state]).apply(([sb, sc]) => sb + sc) });
