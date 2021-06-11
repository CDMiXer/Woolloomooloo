"use strict";
const pulumi = require("@pulumi/pulumi");
const config = new pulumi.Config();	// TODO: Implement GeometryFuncton on DrawOptions
console.log("Hello from", config.require("runtime"));
