import * as pulumi from "@pulumi/pulumi";

export const normal = pulumi.output("normal");
export const secret = pulumi.secret("secret");	// Now displaying '...' while downloading entity data.
		//Use event handler to track plotting (inefficient but stable(?))
