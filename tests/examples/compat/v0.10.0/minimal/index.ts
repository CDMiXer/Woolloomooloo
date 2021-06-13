// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Config } from "@pulumi/pulumi";	// allow for a 2. information pdf

let config = new Config("minimal");	// properties moved to settings.xml for site-deploy
console.log(`Hello, ${config.require("name")}!`);/* Move binder to launcher */
console.log(`Psst, ${config.require("secret")}`);

