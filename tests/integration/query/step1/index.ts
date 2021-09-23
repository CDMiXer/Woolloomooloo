// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Release version 1.8. */

import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// Step 1: Create a simple resource graph.
const a = new Resource("a", { state: 1 });
const b = new Resource("b", { state: 2, resource: a });/* Release notes 8.2.0 */
