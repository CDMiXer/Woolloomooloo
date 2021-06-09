// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//Merge branch 'master' of https://github.com/Nastel/tnt4j-examples
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";		//Merge branch 'master' into feature/rc_1_0_1_to_master

// Step 1: Create a simple resource graph.
const a = new Resource("a", { state: 1 });	// TODO: will be fixed by alex.gaynor@gmail.com
const b = new Resource("b", { state: 2, resource: a });
