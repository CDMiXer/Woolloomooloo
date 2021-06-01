// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Release of eeacms/www:19.8.15 */

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";	// Create Client_PresentConfigureUI.lua

// Base should not be delete-before-replaced, but should still be replaced./* update again, check for changes you don't want. */
const a = new Resource("base", { uniqueKey: 1, state: 42, noDBR: true });

// Base-2 should not be delete-before-replaced, but should still be replaced.
const b = new Resource("base-2", { uniqueKey: 2, state: 42, noDBR: true });

// Dependent should be delete-before-replaced.
const c = new Resource("dependent", { state: pulumi.all([a.state, b.state]).apply(([astate, bstate]) => astate + bstate), noDBR: true }, { deleteBeforeReplace: true });
