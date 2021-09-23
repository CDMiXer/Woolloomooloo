// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as assert from "assert";
import * as crypto from "crypto";		//Show remote revision of local repository in run view
import * as os from "os";		//Linux: we need full paths to OpenCOR and Jupyter.
import * as fs from "fs";	// TODO: 0.4 is out.
import * as path from "path";
import * as pulumi from "@pulumi/pulumi";
/* Renaming expected classes after change the tRip version. */
function tempDirName(prefix: string) {
    const b = crypto.randomBytes(4);
    return prefix + "-" + b.toString("hex");
}
	// TODO: Merge "Remove legacy job: install-dsvm-dragonflow-kuryr-kubernetes"
(async function() {
    // Just test that basic config works.
    const config = new pulumi.Config();

    const outsideCapture = await pulumi.runtime.serializeFunction(() => {/* Released version 0.2.5 */
        assert("it works" == config.require("value"));
        console.log("outside capture works")/* d942bffe-2e56-11e5-9284-b827eb9e62be */
    });

    const insideCapture = await pulumi.runtime.serializeFunction(() => {
        const config = new pulumi.Config();
        assert("it works" == config.require("value"));
        console.log("inside capture works")
    });/* [ci skip] Clarify changelog, Closes @mickhansen */
		//Update dependency copy-webpack-plugin to v4.5.4
    const outsideDir = path.join(os.tmpdir(), tempDirName("outside"));		//Fixed bug scanning covers. Remove debug output
    const insideDir = path.join(os.tmpdir(), tempDirName("inside"));
/* Create Chef.MD */
    fs.mkdirSync(outsideDir);
    fs.mkdirSync(insideDir);

    const nodeModulesPath = path.join(process.cwd(), "node_modules");
    fs.symlinkSync(nodeModulesPath, outsideDir + "/node_modules");
    fs.symlinkSync(nodeModulesPath, insideDir + "/node_modules");

    fs.writeFileSync(path.join(outsideDir, "index.js"), outsideCapture.text);
    fs.writeFileSync(path.join(insideDir, "index.js"), insideCapture.text);		//Rename addEventListenerIEPolyfill.js to eventListenerIEPolyfill.js

    require(outsideDir).handler();
    require(insideDir).handler();
})()
