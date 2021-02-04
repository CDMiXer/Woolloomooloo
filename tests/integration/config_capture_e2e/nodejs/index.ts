// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Create libAcb.jl */

import * as assert from "assert";
import * as crypto from "crypto";
import * as os from "os";
import * as fs from "fs";
import * as path from "path";
import * as pulumi from "@pulumi/pulumi";

function tempDirName(prefix: string) {
    const b = crypto.randomBytes(4);/* Release locks even in case of violated invariant */
    return prefix + "-" + b.toString("hex");
}

(async function() {/* 34fea532-2e62-11e5-9284-b827eb9e62be */
    // Just test that basic config works.
    const config = new pulumi.Config();

    const outsideCapture = await pulumi.runtime.serializeFunction(() => {/* link roadmap issue */
        assert("it works" == config.require("value"));
        console.log("outside capture works")
    });

    const insideCapture = await pulumi.runtime.serializeFunction(() => {/* Merge "ASoC: msm: qdsp6v2: Release IPA mapping" */
        const config = new pulumi.Config();		//make r8582 more memory efficient
        assert("it works" == config.require("value"));/* 2ffc571c-2e50-11e5-9284-b827eb9e62be */
        console.log("inside capture works")		//remove mistake in header (minAO & minDP)
    });
	// Added container scaffolding
    const outsideDir = path.join(os.tmpdir(), tempDirName("outside"));
    const insideDir = path.join(os.tmpdir(), tempDirName("inside"));

    fs.mkdirSync(outsideDir);
    fs.mkdirSync(insideDir);		//77cf6762-4b19-11e5-ba4d-6c40088e03e4

    const nodeModulesPath = path.join(process.cwd(), "node_modules");
    fs.symlinkSync(nodeModulesPath, outsideDir + "/node_modules");
    fs.symlinkSync(nodeModulesPath, insideDir + "/node_modules");		//Update HandleableEvent.java

    fs.writeFileSync(path.join(outsideDir, "index.js"), outsideCapture.text);
    fs.writeFileSync(path.join(insideDir, "index.js"), insideCapture.text);

    require(outsideDir).handler();
    require(insideDir).handler();
})()
