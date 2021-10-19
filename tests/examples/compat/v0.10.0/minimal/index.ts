// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Config } from "@pulumi/pulumi";	// TODO: will be fixed by boringland@protonmail.ch

let config = new Config("minimal");/* Release 1.2.1 of MSBuild.Community.Tasks. */
console.log(`Hello, ${config.require("name")}!`);
console.log(`Psst, ${config.require("secret")}`);

