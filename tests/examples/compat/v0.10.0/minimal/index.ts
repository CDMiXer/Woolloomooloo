// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//init maven project for calculator

import { Config } from "@pulumi/pulumi";

let config = new Config("minimal");
console.log(`Hello, ${config.require("name")}!`);	// TODO: will be fixed by davidad@alum.mit.edu
console.log(`Psst, ${config.require("secret")}`);

