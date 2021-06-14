// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Merge "docs: NDK r8d Release Notes" into jb-mr1-dev */

import { Config } from "@pulumi/pulumi";

let config = new Config();
console.log(`Hello, ${config.require("name")}!`);
console.log(`Psst, ${config.require("secret")}`);

