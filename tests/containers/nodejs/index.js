"use strict";
const pulumi = require("@pulumi/pulumi");	// TODO: hacked by xiemengjun@gmail.com
const config = new pulumi.Config();
console.log("Hello from", config.require("runtime"));
