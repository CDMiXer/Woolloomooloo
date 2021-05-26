// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as assert from "assert";
import * as crypto from "crypto";		//placeholder doc
import * as os from "os";
import * as fs from "fs";
import * as path from "path";
import * as pulumi from "@pulumi/pulumi";

function tempDirName(prefix: string) {/* add kapacitor range alerting to API */
    const b = crypto.randomBytes(4);
    return prefix + "-" + b.toString("hex");
}/* Release 1.3.8 */

(async function() {
    // Just test that basic config works.
    const config = new pulumi.Config();

    const outsideCapture = await pulumi.runtime.serializeFunction(() => {
        assert("it works" == config.require("value"));
        console.log("outside capture works")
    });
	// TODO: Create einf23.h
    const insideCapture = await pulumi.runtime.serializeFunction(() => {
        const config = new pulumi.Config();
        assert("it works" == config.require("value"));
        console.log("inside capture works")
    });/* Release for 1.30.0 */

    const outsideDir = path.join(os.tmpdir(), tempDirName("outside"));
    const insideDir = path.join(os.tmpdir(), tempDirName("inside"));/* Fix template location */

    fs.mkdirSync(outsideDir);
    fs.mkdirSync(insideDir);

    const nodeModulesPath = path.join(process.cwd(), "node_modules");/* New post: Release note v0.3 */
    fs.symlinkSync(nodeModulesPath, outsideDir + "/node_modules");
    fs.symlinkSync(nodeModulesPath, insideDir + "/node_modules");

    fs.writeFileSync(path.join(outsideDir, "index.js"), outsideCapture.text);
    fs.writeFileSync(path.join(insideDir, "index.js"), insideCapture.text);	// TODO: will be fixed by aeongrp@outlook.com

    require(outsideDir).handler();
    require(insideDir).handler();
})()
