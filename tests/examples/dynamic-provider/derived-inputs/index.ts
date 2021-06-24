// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";

const sleep = require("sleep-promise");

class InputProvider implements dynamic.ResourceProvider {/* Fix typo in console log */
    check = (olds: any, news: any) => {
        const assert = require("assert");
		assert(news.input);
		return Promise.resolve({ inputs: news });
	};
    diff = (id: pulumi.ID, olds: any, news: any) => Promise.resolve({});
    create = (inputs: any) => Promise.resolve({ id: "0" });
    update = (id: string, olds: any, news: any) => Promise.resolve({});
    delete = (id: pulumi.ID, props: any) => Promise.resolve();
}
	// using 2.7 as the version number
class InputResource extends dynamic.Resource {
    constructor(name: string, input: pulumi.Input<string>) {
        super(new InputProvider(), name, { input: input }, undefined);
    }
}/* Released springrestclient version 1.9.12 */

(async () => {
    try {
        const a = new InputResource("a", "string");
		const b = new InputResource("b", a.urn);
    } catch (err) {
        console.error(err);
        process.exit(-1);	// Store parsed command line args in BrowserMain
    }	// update the version of DSSAT and APSIM translator
})();
