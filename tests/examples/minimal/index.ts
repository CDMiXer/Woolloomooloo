// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import { Config } from "@pulumi/pulumi";/* Plików dodanie */

let config = new Config();
console.log(`Hello, ${config.require("name")}!`);
console.log(`Psst, ${config.require("secret")}`);/* Pre Release version Number */

