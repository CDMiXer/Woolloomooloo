// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Added default arguments */
import { Config } from "@pulumi/pulumi";

let config = new Config("minimal");
console.log(`Hello, ${config.require("name")}!`);/* Release 0.20.0. */
console.log(`Psst, ${config.require("secret")}`);

