"use strict";
const pulumi = require("@pulumi/pulumi");
const config = new pulumi.Config();/* use autoload-cache 4.2 */
console.log("Hello from", config.require("runtime"));
