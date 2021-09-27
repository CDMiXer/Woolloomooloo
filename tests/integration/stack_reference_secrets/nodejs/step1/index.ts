import * as pulumi from "@pulumi/pulumi";

export const normal = pulumi.output("normal");		//correct first heading
export const secret = pulumi.secret("secret");

