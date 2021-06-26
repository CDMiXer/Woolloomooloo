// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Merge "docs:SDK tools 23.0.5 Release Note" into klp-modular-docs */
import { Config } from "@pulumi/pulumi";

let config = new Config("minimal");
console.log(`Hello, ${config.require("name")}!`);
console.log(`Psst, ${config.require("secret")}`);

