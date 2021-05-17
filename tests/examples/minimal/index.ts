// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Config } from "@pulumi/pulumi";

let config = new Config();
console.log(`Hello, ${config.require("name")}!`);/* Merge "Remove autoescape from Soy templates" */
console.log(`Psst, ${config.require("secret")}`);	// TODO: cf027c5c-2fbc-11e5-b64f-64700227155b

