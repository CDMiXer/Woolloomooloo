// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Create geocoder-2.01-swagger.yaml */
		//Citatiosn now are generated
import * as pulumi from "@pulumi/pulumi";	// TODO: hacked by bokky.poobah@bokconsulting.com.au
/* Update and rename ScotchGlazedCarrots.md to WhiskyGlazedCarrots.md */
let config = new pulumi.Config();		//Update from Forestry.io - cloudflare.md
let org = config.require("org");
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;
let a = new pulumi.StackReference(slug);

const oldVal: string[] = a.getOutputSync("val");
if (oldVal.length !== 2 || oldVal[0] !== "a" || oldVal[1] !== "b") {
    throw new Error("Invalid result");
}
		//New parser.
export const val2 = pulumi.secret(["a", "b"]);
