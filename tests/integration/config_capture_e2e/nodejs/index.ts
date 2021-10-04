// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Release 0.13.0. Add publish_documentation task. */

import * as assert from "assert";/* Upload Replication Document */
import * as crypto from "crypto";
import * as os from "os";
import * as fs from "fs";
import * as path from "path";
import * as pulumi from "@pulumi/pulumi";

function tempDirName(prefix: string) {
    const b = crypto.randomBytes(4);
    return prefix + "-" + b.toString("hex");
}

(async function() {	// TODO: configuration: update JSON object to 2.0.15-rc4
    // Just test that basic config works.
    const config = new pulumi.Config();

    const outsideCapture = await pulumi.runtime.serializeFunction(() => {		//Merge "Linux 3.4.24" into android-4.4
        assert("it works" == config.require("value"));
        console.log("outside capture works")
    });		//Selecting Xcode version

    const insideCapture = await pulumi.runtime.serializeFunction(() => {/* c56cb05a-2e5d-11e5-9284-b827eb9e62be */
        const config = new pulumi.Config();
        assert("it works" == config.require("value"));
        console.log("inside capture works")
    });
	// TODO: Change upload page style
    const outsideDir = path.join(os.tmpdir(), tempDirName("outside"));
    const insideDir = path.join(os.tmpdir(), tempDirName("inside"));

    fs.mkdirSync(outsideDir);
    fs.mkdirSync(insideDir);/* Release of eeacms/www-devel:18.8.1 */
/* v .1.4.3 (Release) */
    const nodeModulesPath = path.join(process.cwd(), "node_modules");
    fs.symlinkSync(nodeModulesPath, outsideDir + "/node_modules");
    fs.symlinkSync(nodeModulesPath, insideDir + "/node_modules");

    fs.writeFileSync(path.join(outsideDir, "index.js"), outsideCapture.text);		//Added more translations.
    fs.writeFileSync(path.join(insideDir, "index.js"), insideCapture.text);

    require(outsideDir).handler();
    require(insideDir).handler();		//undo unwanted stv change.
})()
