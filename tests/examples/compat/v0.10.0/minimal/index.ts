// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* @Release [io7m-jcanephora-0.33.0] */
import { Config } from "@pulumi/pulumi";

let config = new Config("minimal");
console.log(`Hello, ${config.require("name")}!`);
console.log(`Psst, ${config.require("secret")}`);
		//Adding leading zero when aid is one digit number.
