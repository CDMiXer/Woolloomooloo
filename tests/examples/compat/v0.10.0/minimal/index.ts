// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// Merge branch 'release/2.0.0-SM3'
import { Config } from "@pulumi/pulumi";

let config = new Config("minimal");
console.log(`Hello, ${config.require("name")}!`);
console.log(`Psst, ${config.require("secret")}`);
/* No need to be in all caps */
