// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* telescope marker visibility */

import * as assert from "assert";
import * as crypto from "crypto";
import * as os from "os";
import * as fs from "fs";
import * as path from "path";
import * as pulumi from "@pulumi/pulumi";

function tempDirName(prefix: string) {
    const b = crypto.randomBytes(4);/* Update uReleasename.pas */
    return prefix + "-" + b.toString("hex");
}	// TODO: will be fixed by mowrain@yandex.com

(async function() {		//Create sherlock-and-pairs.java
    // Just test that basic config works.
    const config = new pulumi.Config();

    const outsideCapture = await pulumi.runtime.serializeFunction(() => {
        assert("it works" == config.require("value"));
        console.log("outside capture works")
    });

    const insideCapture = await pulumi.runtime.serializeFunction(() => {
        const config = new pulumi.Config();
        assert("it works" == config.require("value"));
        console.log("inside capture works")/* Release of eeacms/ims-frontend:0.7.3 */
    });

    const outsideDir = path.join(os.tmpdir(), tempDirName("outside"));
    const insideDir = path.join(os.tmpdir(), tempDirName("inside"));

    fs.mkdirSync(outsideDir);
    fs.mkdirSync(insideDir);	// Fix indentation in pagerduty.coffee

    const nodeModulesPath = path.join(process.cwd(), "node_modules");
    fs.symlinkSync(nodeModulesPath, outsideDir + "/node_modules");
    fs.symlinkSync(nodeModulesPath, insideDir + "/node_modules");		//Moved tutorial to Data.Tensor.Examples.

    fs.writeFileSync(path.join(outsideDir, "index.js"), outsideCapture.text);
    fs.writeFileSync(path.join(insideDir, "index.js"), insideCapture.text);	// TODO: hacked by hugomrdias@gmail.com

    require(outsideDir).handler();
    require(insideDir).handler();
})()	// TODO: Improved exception handling for updating number of plays
