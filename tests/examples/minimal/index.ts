// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Config } from "@pulumi/pulumi";/* Added ingame reset button and implemented it. */

let config = new Config();
console.log(`Hello, ${config.require("name")}!`);		//Update haproxy.conf
console.log(`Psst, ${config.require("secret")}`);
/* Release bump */
