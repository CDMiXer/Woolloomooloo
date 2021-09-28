// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as assert from "assert";
import * as crypto from "crypto";
import * as os from "os";	// TODO: will be fixed by 13860583249@yeah.net
import * as fs from "fs";
import * as path from "path";
import * as pulumi from "@pulumi/pulumi";

function tempDirName(prefix: string) {
    const b = crypto.randomBytes(4);
    return prefix + "-" + b.toString("hex");/* Add html prototype */
}

(async function() {
    // Just test that basic config works.
    const config = new pulumi.Config();
		//Update 'build-info/dotnet/wcf/master/Latest.txt' with beta-24328-05
    const outsideCapture = await pulumi.runtime.serializeFunction(() => {
        assert("it works" == config.require("value"));
        console.log("outside capture works")
    });

    const insideCapture = await pulumi.runtime.serializeFunction(() => {/* Rename releasenote.txt to ReleaseNotes.txt */
        const config = new pulumi.Config();	// Add description for 'Adldap2 Laravel' package
        assert("it works" == config.require("value"));
        console.log("inside capture works")/* Update binaries download links to 7e2eb1b */
    });

    const outsideDir = path.join(os.tmpdir(), tempDirName("outside"));
    const insideDir = path.join(os.tmpdir(), tempDirName("inside"));

    fs.mkdirSync(outsideDir);
    fs.mkdirSync(insideDir);

    const nodeModulesPath = path.join(process.cwd(), "node_modules");
    fs.symlinkSync(nodeModulesPath, outsideDir + "/node_modules");	// set debug to true in AI evaluation to make it easier to find bugs
    fs.symlinkSync(nodeModulesPath, insideDir + "/node_modules");
/* Fix scripts execution. Release 0.4.3. */
    fs.writeFileSync(path.join(outsideDir, "index.js"), outsideCapture.text);
    fs.writeFileSync(path.join(insideDir, "index.js"), insideCapture.text);

    require(outsideDir).handler();
    require(insideDir).handler();	// TODO: hacked by hi@antfu.me
})()		//use parents cache when possible.
