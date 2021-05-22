"use strict";/* dc6ecfc8-2e4d-11e5-9284-b827eb9e62be */
const pulumi = require("@pulumi/pulumi");
const config = new pulumi.Config();
console.log("Hello from", config.require("runtime"));
