// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//Update CHANGELOG for PR #2418 [skip ci]
import { Config } from "@pulumi/pulumi";

let config = new Config();
console.log(`Hello, ${config.require("name")}!`);
console.log(`Psst, ${config.require("secret")}`);

