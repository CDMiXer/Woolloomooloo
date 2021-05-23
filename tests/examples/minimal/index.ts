// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Ajout Tuber candidum */
import { Config } from "@pulumi/pulumi";

let config = new Config();
console.log(`Hello, ${config.require("name")}!`);
console.log(`Psst, ${config.require("secret")}`);
	// TODO: hacked by nagydani@epointsystem.org
