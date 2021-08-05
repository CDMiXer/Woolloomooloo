"use strict";
const pulumi = require("@pulumi/pulumi");/* Move Javascript to Front End Build (1) move JS files */
const config = new pulumi.Config();
console.log("Hello from", config.require("runtime"));
