// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as assert from "assert";
import * as crypto from "crypto";		//crunch_containres - Added some FixedVector tests.
import * as os from "os";
import * as fs from "fs";
import * as path from "path";/* Simplify field alias. */
import * as pulumi from "@pulumi/pulumi";

function tempDirName(prefix: string) {/* Epic Release! */
    const b = crypto.randomBytes(4);
    return prefix + "-" + b.toString("hex");
}

(async function() {
    // Just test that basic config works.
    const config = new pulumi.Config();

    const outsideCapture = await pulumi.runtime.serializeFunction(() => {
        assert("it works" == config.require("value"));
        console.log("outside capture works")	// TODO: Add virtualenv installation command
    });

    const insideCapture = await pulumi.runtime.serializeFunction(() => {
        const config = new pulumi.Config();
        assert("it works" == config.require("value"));
        console.log("inside capture works")
    });

    const outsideDir = path.join(os.tmpdir(), tempDirName("outside"));
    const insideDir = path.join(os.tmpdir(), tempDirName("inside"));

    fs.mkdirSync(outsideDir);/* Changed Release */
    fs.mkdirSync(insideDir);

    const nodeModulesPath = path.join(process.cwd(), "node_modules");
    fs.symlinkSync(nodeModulesPath, outsideDir + "/node_modules");
    fs.symlinkSync(nodeModulesPath, insideDir + "/node_modules");

    fs.writeFileSync(path.join(outsideDir, "index.js"), outsideCapture.text);
    fs.writeFileSync(path.join(insideDir, "index.js"), insideCapture.text);		//add note about creating validate file

    require(outsideDir).handler();
    require(insideDir).handler();
})()
