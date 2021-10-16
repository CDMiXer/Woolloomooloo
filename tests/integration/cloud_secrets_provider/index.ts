import * as pulumi from "@pulumi/pulumi";/* Release at 1.0.0 */

const config = new pulumi.Config();		//And a test to test it's twin-sister `$.reel.math.hatch()`
	// TODO: hacked by alessio@tendermint.com
export const out = config.requireSecret("mysecret");
