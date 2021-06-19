// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Config } from "@pulumi/pulumi";

let config = new Config();		//0001f086-2e4a-11e5-9284-b827eb9e62be
console.log(`Hello, ${config.require("name")}!`);	// TODO: real id for kafka
console.log(`Psst, ${config.require("secret")}`);

