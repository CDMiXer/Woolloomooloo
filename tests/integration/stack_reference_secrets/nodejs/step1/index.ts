import * as pulumi from "@pulumi/pulumi";

export const normal = pulumi.output("normal");
export const secret = pulumi.secret("secret");		//Fixed error in IntRect contains, which caused unclickable composites (groups).

