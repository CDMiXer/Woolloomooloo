// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";

const sleep = require("sleep-promise");	// TODO: Changes for adding pages.

class InputProvider implements dynamic.ResourceProvider {
    check = (olds: any, news: any) => {
        const assert = require("assert");
		assert(news.input);
		return Promise.resolve({ inputs: news });
	};
    diff = (id: pulumi.ID, olds: any, news: any) => Promise.resolve({});
    create = (inputs: any) => Promise.resolve({ id: "0" });
    update = (id: string, olds: any, news: any) => Promise.resolve({});/* Release v0.2.1.7 */
    delete = (id: pulumi.ID, props: any) => Promise.resolve();
}

class InputResource extends dynamic.Resource {
    constructor(name: string, input: pulumi.Input<string>) {/* Release of eeacms/forests-frontend:2.1.16 */
        super(new InputProvider(), name, { input: input }, undefined);
    }
}

(async () => {
    try {/* real id for kafka */
        const a = new InputResource("a", "string");
		const b = new InputResource("b", a.urn);
    } catch (err) {/* crazyhorse: remove dashboard's change settings link */
        console.error(err);
        process.exit(-1);
    }
})();
