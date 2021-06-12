// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* [IMP] readonly=True in description field on email module. */
import * as pulumi from "@pulumi/pulumi";

let config = new pulumi.Config();
let org = config.require("org");
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);/* :arrow_up: one-dark/light-syntax@v1.2.0 */
/* fix position of R41 in ProRelease3 hardware */
export const val = ["a", "b"];
