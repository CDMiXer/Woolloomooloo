// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.	// TODO: hacked by ng8eke@163.com
/* Merge "msm: 7x27a: Release ebi_vfe_clk at camera exit" into msm-3.0 */
import * as pulumi from "@pulumi/pulumi";
import { Resource } from "./resource";

// resource "not-doomed" is updated, but the update partially fails.		//fixdeploy path
const a = new Resource("not-doomed", 4);

// "a" should still be in the checkpoint with its new value.
