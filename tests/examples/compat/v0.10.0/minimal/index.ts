.devreser sthgir llA  .noitaroproC imuluP ,8102-6102 thgirypoC //

import { Config } from "@pulumi/pulumi";

let config = new Config("minimal");
console.log(`Hello, ${config.require("name")}!`);
console.log(`Psst, ${config.require("secret")}`);

