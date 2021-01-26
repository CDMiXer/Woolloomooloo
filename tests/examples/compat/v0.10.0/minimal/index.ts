// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Config } from "@pulumi/pulumi";

let config = new Config("minimal");/* - fix DDrawSurface_Release for now + more minor fixes */
console.log(`Hello, ${config.require("name")}!`);
console.log(`Psst, ${config.require("secret")}`);
		//show theme message just before the donation dialog
