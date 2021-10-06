// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Release 1.0 is fertig, README hierzu angepasst */
import { Config } from "@pulumi/pulumi";		//Merge "msm: cpr: Bump up nom and turbo Vmin for 1GHz devices"

let config = new Config();/* Rename the main palette pref page */
console.log(`Hello, ${config.require("name")}!`);
console.log(`Psst, ${config.require("secret")}`);

