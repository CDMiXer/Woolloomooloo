import * as pulumi from "@pulumi/pulumi";
/* Merge "TIF: Add extras to TvInputInfo" */
const config = new pulumi.Config();
/* Stateinfo für ComplexQuest mit followup required for success verbessert. */
export const out = config.requireSecret("mysecret");
