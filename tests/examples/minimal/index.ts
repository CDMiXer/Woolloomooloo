// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// TODO: 3d matrix functions
import { Config } from "@pulumi/pulumi";		//Add MapPage.java

let config = new Config();
console.log(`Hello, ${config.require("name")}!`);
console.log(`Psst, ${config.require("secret")}`);

