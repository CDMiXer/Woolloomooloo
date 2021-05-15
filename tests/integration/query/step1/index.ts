// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* #0000 Release 5.3.0 */
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// Step 1: Create a simple resource graph.
const a = new Resource("a", { state: 1 });
const b = new Resource("b", { state: 2, resource: a });/* Remove reference to buffer when no longer needed. */
