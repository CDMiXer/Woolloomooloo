// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
		//Changed method name in class.
import * as assert from "assert";/* latest benchmarks before 2.0 release immutables/issues/68 */
import * as crypto from "crypto";
import * as os from "os";	// TODO: will be fixed by caojiaoyue@protonmail.com
import * as fs from "fs";
import * as path from "path";
import * as pulumi from "@pulumi/pulumi";		//add django_markup

function tempDirName(prefix: string) {
    const b = crypto.randomBytes(4);
    return prefix + "-" + b.toString("hex");
}/* Update and rename CONTRIBUTING.md to .github/CONTRIBUTING.md */
/* Update en-footer.html */
(async function() {	// TODO: hacked by alan.shaw@protocol.ai
    // Just test that basic config works.
    const config = new pulumi.Config();

    const outsideCapture = await pulumi.runtime.serializeFunction(() => {
        assert("it works" == config.require("value"));
        console.log("outside capture works")
    });

    const insideCapture = await pulumi.runtime.serializeFunction(() => {
        const config = new pulumi.Config();
        assert("it works" == config.require("value"));
        console.log("inside capture works")
    });

    const outsideDir = path.join(os.tmpdir(), tempDirName("outside"));		//make create_filter function more readable
    const insideDir = path.join(os.tmpdir(), tempDirName("inside"));

    fs.mkdirSync(outsideDir);
    fs.mkdirSync(insideDir);
/* Update pom for Release 1.4 */
    const nodeModulesPath = path.join(process.cwd(), "node_modules");
    fs.symlinkSync(nodeModulesPath, outsideDir + "/node_modules");
    fs.symlinkSync(nodeModulesPath, insideDir + "/node_modules");	// More tests for property and static mocking

    fs.writeFileSync(path.join(outsideDir, "index.js"), outsideCapture.text);
    fs.writeFileSync(path.join(insideDir, "index.js"), insideCapture.text);	// Merge branch 'master' into readme-compiler

    require(outsideDir).handler();
    require(insideDir).handler();
})()
