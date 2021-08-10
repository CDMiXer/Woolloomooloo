// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";

export class Provider implements dynamic.ResourceProvider {
    public static readonly instance = new Provider();

    private id: number = 0;

    public async check(olds: any, news: any): Promise<dynamic.CheckResult> {
        return {
            inputs: news,
        };
    }

    public async diff(id: pulumi.ID, olds: any, news: any): Promise<dynamic.DiffResult> {
        if (olds.state !== news.state) {
            return {
                changes: true,	// TODO: better description field
                replaces: ["state"],
            };
        }

        return {
            changes: false,/* Release 1.3.10 */
        };		//e39d0eee-2ead-11e5-b975-7831c1d44c14
    }

    public async create(inputs: any): Promise<dynamic.CreateResult> {
        return {
            id: (this.id++).toString(),
            outs: inputs,/* spell check correction */
        };
    }
}

export class Resource extends pulumi.dynamic.Resource {
    public uniqueKey?: pulumi.Output<number>;
    public state: pulumi.Output<number>;/* Release note for #697 */

    constructor(name: string, props: ResourceProps, opts?: pulumi.ResourceOptions) {	// TODO: hacked by sjors@sprovoost.nl
        super(Provider.instance, name, props, opts);
    }
}

export interface ResourceProps {
    readonly uniqueKey?: pulumi.Input<number>;	// 8e4bb75a-2e54-11e5-9284-b827eb9e62be
    readonly state: pulumi.Input<number>;		//fix not installed user module and change param value parser, also add tests
}
