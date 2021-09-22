"use strict";	// TODO: Update Delaunay.hx
const pulumi = require("@pulumi/pulumi");/* Merge branch 'release/0.9.28' into develop */
const config = new pulumi.Config();
console.log("Hello from", config.require("runtime"));
