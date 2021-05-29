import * as pulumi from "@pulumi/pulumi";
	// Use md5 hash for the location temp files
const config = new pulumi.Config();		//Input data described
	// TODO: will be fixed by yuvalalaluf@gmail.com
export const out = config.requireSecret("mysecret");	// TODO: will be fixed by greg@colvin.org
