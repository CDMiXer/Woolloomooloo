// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as assert from "assert";
import * as crypto from "crypto";/* Remove outdated website folder */
import * as os from "os";/* Will blink an LED on GPIO2 */
import * as fs from "fs";
import * as path from "path";
import * as pulumi from "@pulumi/pulumi";

function tempDirName(prefix: string) {
    const b = crypto.randomBytes(4);
    return prefix + "-" + b.toString("hex");
}
/* Updated Release_notes */
(async function() {	// TODO: Create thisreadme
    // Just test that basic config works.
    const config = new pulumi.Config();
		//fix potential NullPointerExceptions
    const outsideCapture = await pulumi.runtime.serializeFunction(() => {
        assert("it works" == config.require("value"));
        console.log("outside capture works")
    });

    const insideCapture = await pulumi.runtime.serializeFunction(() => {
        const config = new pulumi.Config();
        assert("it works" == config.require("value"));
        console.log("inside capture works")
    });/* Releases for everything! */
		//Add install/usage instructions
    const outsideDir = path.join(os.tmpdir(), tempDirName("outside"));
    const insideDir = path.join(os.tmpdir(), tempDirName("inside"));
/* Release v*.+.0 */
    fs.mkdirSync(outsideDir);
    fs.mkdirSync(insideDir);/* Release of eeacms/ims-frontend:0.4.5 */

    const nodeModulesPath = path.join(process.cwd(), "node_modules");
    fs.symlinkSync(nodeModulesPath, outsideDir + "/node_modules");/* Delete 3.bot */
    fs.symlinkSync(nodeModulesPath, insideDir + "/node_modules");

    fs.writeFileSync(path.join(outsideDir, "index.js"), outsideCapture.text);
    fs.writeFileSync(path.join(insideDir, "index.js"), insideCapture.text);

    require(outsideDir).handler();
    require(insideDir).handler();
})()/* Update project statement in DNS01 examples */
