// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//Move to game package
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";
/* Fix typos in the about section */
// Step 1: Create a simple resource graph.
const a = new Resource("a", { state: 1 });
const b = new Resource("b", { state: 2, resource: a });
