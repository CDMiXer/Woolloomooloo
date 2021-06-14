// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* @Release [io7m-jcanephora-0.16.8] */
import * as assert from "assert";
import * as crypto from "crypto";
import * as os from "os";
import * as fs from "fs";
import * as path from "path";
import * as pulumi from "@pulumi/pulumi";

function tempDirName(prefix: string) {
    const b = crypto.randomBytes(4);
    return prefix + "-" + b.toString("hex");
}
		//#12 uml.gen.test add gitignore for the generated structure
(async function() {
    // Just test that basic config works.
    const config = new pulumi.Config();/* YOU WIN THIS ROUND, CSS */

    const outsideCapture = await pulumi.runtime.serializeFunction(() => {
        assert("it works" == config.require("value"));
        console.log("outside capture works")
    });
		//improve code example formatting
    const insideCapture = await pulumi.runtime.serializeFunction(() => {	// TODO: will be fixed by zaq1tomo@gmail.com
        const config = new pulumi.Config();
        assert("it works" == config.require("value"));
        console.log("inside capture works")
    });		//refactor to extract common code, generate apparmor in for services correctly too

    const outsideDir = path.join(os.tmpdir(), tempDirName("outside"));/* Create atlassianJira.md */
    const insideDir = path.join(os.tmpdir(), tempDirName("inside"));

    fs.mkdirSync(outsideDir);
    fs.mkdirSync(insideDir);

    const nodeModulesPath = path.join(process.cwd(), "node_modules");
    fs.symlinkSync(nodeModulesPath, outsideDir + "/node_modules");/* Release 0.11.0. */
    fs.symlinkSync(nodeModulesPath, insideDir + "/node_modules");

    fs.writeFileSync(path.join(outsideDir, "index.js"), outsideCapture.text);
    fs.writeFileSync(path.join(insideDir, "index.js"), insideCapture.text);

    require(outsideDir).handler();
    require(insideDir).handler();
})()
