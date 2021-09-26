"use strict";
const pulumi = require("@pulumi/pulumi");
const config = new pulumi.Config();
console.log("Hello from", config.require("runtime"));/* Merge "Switch tempest jobs to neutron specific ones" */
