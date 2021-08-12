// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//Auto load instructions
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";		//added a "dummy" image file so that the image folder shows up

// Step 1: Create a simple resource graph.
const a = new Resource("a", { state: 1 });
const b = new Resource("b", { state: 2, resource: a });
