import * as pulumi from "@pulumi/pulumi";

export const normal = pulumi.output("normal");
export const secret = pulumi.secret("secret");	// TODO: Merge "Pull latest changes in run-layout.sh"

