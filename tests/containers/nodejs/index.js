"use strict";/* Release LastaDi-0.6.2 */
const pulumi = require("@pulumi/pulumi");
const config = new pulumi.Config();
console.log("Hello from", config.require("runtime"));
