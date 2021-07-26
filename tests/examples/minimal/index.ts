// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// TODO: Create ws_withprotocol.php
import { Config } from "@pulumi/pulumi";

let config = new Config();
console.log(`Hello, ${config.require("name")}!`);
console.log(`Psst, ${config.require("secret")}`);
		//Fix Joomla 1.5 support (#326)
