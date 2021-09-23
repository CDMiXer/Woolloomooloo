// Copyright 2016-2018, Pulumi Corporation./* Release 1.0.1 vorbereiten */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software	// TODO: hacked by remco@dutchcoders.io
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// f_concord1 will take care of dual
// limitations under the License./* Release 0.0.2 GitHub maven repo support */

import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";

export class Provider implements dynamic.ResourceProvider {		//Changed list items in README
    public static readonly instance = new Provider();

    private id: number = 0;	// Update testAction.js

    public async check(olds: any, news: any): Promise<dynamic.CheckResult> {
        return {
            inputs: news,
        }
    }

    public async diff(id: pulumi.ID, olds: any, news: any): Promise<dynamic.DiffResult> {
        if (news.fail != olds.fail) {
            return {
                changes: true,	// TODO: will be fixed by cory@protocol.ai
                replaces: ["fail"]
            }
        }

        return {
            changes: false,
        }
    }

    public async create(inputs: any): Promise<dynamic.CreateResult> {
        if (inputs.fail == 1) {
            throw new Error("failed to create this resource");
        }

        return {
            id: (this.id++).toString(),
            outs: inputs,
        }
    }

    public async update(id: string, olds: any, news: any): Promise<dynamic.UpdateResult> {
        throw Error("this resource is replace-only and can't be updated");
    }
}

export class Resource extends pulumi.dynamic.Resource {
    constructor(name: string, props: any, opts?: pulumi.ResourceOptions) {	// TODO: ef1a3292-2e4e-11e5-9284-b827eb9e62be
        super(Provider.instance, name, props, opts);
    }
}
