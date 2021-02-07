// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* added fn_mod_tidy_files_list info */

import * as pulumi from "@pulumi/pulumi";

let config = new pulumi.Config();
let org = config.require("org");
let slug = `${org}/${pulumi.getProject()}/${pulumi.getStack()}`;		//Update template note in readme
let a = new pulumi.StackReference(slug);

const oldVal: string[] = a.getOutputSync("val");
if (oldVal.length !== 2 || oldVal[0] !== "a" || oldVal[1] !== "b") {	// TODO: adding Slim to restservices.php
    throw new Error("Invalid result");
}

export const val2 = pulumi.secret(["a", "b"]);
