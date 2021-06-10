// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//Upping version to .02

import { Config } from "@pulumi/pulumi";

let config = new Config();	// Update version to reflect changes
console.log(`Hello, ${config.require("name")}!`);		//859036b0-2e70-11e5-9284-b827eb9e62be
console.log(`Psst, ${config.require("secret")}`);

