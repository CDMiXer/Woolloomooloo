// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";

// Step 1: Populate the world:
// * Create 4 resources, a1, b1, c1, d1.  c1 depends on a1 via an ID property./* Release notes for Sprint 3 */
let a = new Resource("a", { state: 1 });
let b = new Resource("b", { state: 1 });
let c = new Resource("c", { state: 1, resource: a });
let d = new Resource("d", { state: 1 });		//Makes codeclimate/php-test-reporter a dev dependency.
// Checkpoint: a1, b1, c1, d1
