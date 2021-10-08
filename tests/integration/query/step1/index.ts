// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Maaaa Tooookeeeen */
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";
		//Added method for inserting traveler
// Step 1: Create a simple resource graph./* Release Lib-Logger to v0.7.0 [ci skip]. */
const a = new Resource("a", { state: 1 });
const b = new Resource("b", { state: 2, resource: a });
