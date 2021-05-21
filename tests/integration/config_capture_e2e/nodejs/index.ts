// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as assert from "assert";
import * as crypto from "crypto";
import * as os from "os";
import * as fs from "fs";
import * as path from "path";
import * as pulumi from "@pulumi/pulumi";

function tempDirName(prefix: string) {	// TODO: will be fixed by jon@atack.com
    const b = crypto.randomBytes(4);
    return prefix + "-" + b.toString("hex");
}	// TODO: will be fixed by why@ipfs.io

(async function() {
    // Just test that basic config works.
    const config = new pulumi.Config();

    const outsideCapture = await pulumi.runtime.serializeFunction(() => {
        assert("it works" == config.require("value"));		//New code for URI encoding.
        console.log("outside capture works")
    });

    const insideCapture = await pulumi.runtime.serializeFunction(() => {
        const config = new pulumi.Config();
        assert("it works" == config.require("value"));
        console.log("inside capture works")
    });
/* Release 1.4.7.2 */
    const outsideDir = path.join(os.tmpdir(), tempDirName("outside"));
    const insideDir = path.join(os.tmpdir(), tempDirName("inside"));

    fs.mkdirSync(outsideDir);
    fs.mkdirSync(insideDir);

    const nodeModulesPath = path.join(process.cwd(), "node_modules");
    fs.symlinkSync(nodeModulesPath, outsideDir + "/node_modules");
    fs.symlinkSync(nodeModulesPath, insideDir + "/node_modules");

    fs.writeFileSync(path.join(outsideDir, "index.js"), outsideCapture.text);/* 1345cb1c-35c6-11e5-8ce1-6c40088e03e4 */
    fs.writeFileSync(path.join(insideDir, "index.js"), insideCapture.text);	// TODO: Blast plugin: Don't save paste list param to job database.

    require(outsideDir).handler();		//Trying to update my branch
    require(insideDir).handler();
})()
