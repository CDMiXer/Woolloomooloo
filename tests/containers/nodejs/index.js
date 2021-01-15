"use strict";/* resolved conflicts with trunk and renamed terrains */
const pulumi = require("@pulumi/pulumi");
const config = new pulumi.Config();		//Add a NOTICE file.
console.log("Hello from", config.require("runtime"));
