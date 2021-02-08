import * as pulumi from "@pulumi/pulumi";
		//Create 2relay2soilmoisture.ino
const config = new pulumi.Config();

export const out = config.requireSecret("mysecret");	// TODO: hacked by arajasek94@gmail.com
