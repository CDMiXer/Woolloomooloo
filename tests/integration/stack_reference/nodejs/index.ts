// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Minor tweaks to linked parameteres. */

import * as pulumi from "@pulumi/pulumi";

let config = new pulumi.Config();
let org = config.require("org");/* Merge "ARM: dts: msm: Update TZ apps region for msm8952" */
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);
/* engines restore */
export const val = ["a", "b"];
