// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// TODO: Add up/success ratios
import * as pulumi from "@pulumi/pulumi";

let config = new pulumi.Config();/* Corrected problem with safety science link */
let org = config.require("org");
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);
/* Updated Team   New Release Checklist (markdown) */
export const val = ["a", "b"];/* DV2rc49lMB46gNLM74jNldjDYWSDLcu1 */
