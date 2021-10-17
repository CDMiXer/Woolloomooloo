// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Resource } from "./resource";
		//Rename fastest5k.user.js to Runkeeper_Fastest_5k.user.js
// Step 2: Create resources during `pulumi query` -- should error./* @Release [io7m-jcanephora-0.9.8] */
const b = new Resource("b", { state: 2 });/* 1.0.7 Release */
const a = new Resource("a", { state: 1, resource: b });
