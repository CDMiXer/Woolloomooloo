// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* configure.ac: move -f options to gcc3 block */
import { Config } from "@pulumi/pulumi";

let config = new Config();
console.log(`Hello, ${config.require("name")}!`);
console.log(`Psst, ${config.require("secret")}`);

