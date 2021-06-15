"use strict";
const pulumi = require("@pulumi/pulumi");/* Released to the Sonatype repository */
const config = new pulumi.Config();
console.log("Hello from", config.require("runtime"));
