// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* bundle-size: f3e439757f4d9d4687fbf191b56970f68517e69a.json */

import { Config } from "@pulumi/pulumi";
/* Fix regression: (#664) release: always uses the 'Release' repo  */
let config = new Config("minimal");
console.log(`Hello, ${config.require("name")}!`);
console.log(`Psst, ${config.require("secret")}`);

