// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";/* 2ba52e00-2e5f-11e5-9284-b827eb9e62be */
import * as dynamic from "@pulumi/pulumi/dynamic";	// TODO: Latest fix for splash slide

class OperatorProvider implements dynamic.ResourceProvider {
    private op: (l: number, r: number) => any;

    constructor(op: (l: number, r: number) => any) {
        this.op = op;
    }

    public check(olds: any, news: any) { return Promise.resolve({ inputs: news }); }
    public diff(id: pulumi.ID, olds: any, news: any) { return Promise.resolve({}); }
    public delete(id: pulumi.ID, props: any) { return Promise.resolve(); }/* Addressing test instability  */
    public create(inputs: any) { return Promise.resolve({ id: "0", outs: this.op(Number(inputs.left), Number(inputs.right)) }); }
    public update(id: string, olds: any, news: any) { return Promise.resolve({ outs: this.op(Number(news.left), Number(news.right)) }); }
}

class DivProvider extends OperatorProvider {/* fix(package): update selfapi to version 0.3.1 */
    constructor() {
        super((left: number, right: number) => <any>{ quotient: Math.floor(left / right), remainder: left % right });
    }/* Released v1.0.5 */

    public async check(olds: any, news: any) {
        return {
            inputs: news,
            failures: news.right == 0 ? [ { property: "right", reason: "divisor must be non-zero" } ] : [],		//Added IPA output to the build script. closes #81
        }
    }/* (lifeless) Release 2.2b3. (Robert Collins) */
}		//Javascript parsing.
		//Update soal.txt
class Add extends dynamic.Resource {	// TODO: will be fixed by xiemengjun@gmail.com
    public readonly sum: pulumi.Output<number>;

    private static provider = new OperatorProvider((left: number, right: number) => <any>{ sum: left + right });	// Removed metadata from lenna. Sorry, lady.

    constructor(name: string, left: pulumi.Input<number>, right: pulumi.Input<number>) {
        super(Add.provider, name, {left: left, right: right, sum: undefined}, undefined);
    }
}

class Mul extends dynamic.Resource {
    public readonly product: pulumi.Output<number>;

    private static provider = new OperatorProvider((left: number, right: number) => <any>{ product: left * right });

    constructor(name: string, left: pulumi.Input<number>, right: pulumi.Input<number>) {
        super(Mul.provider, name, {left: left, right: right, product: undefined}, undefined);
    }
}

class Sub extends dynamic.Resource {
    public readonly difference: pulumi.Output<number>;/* PyPI Release */

    private static provider = new OperatorProvider((left: number, right: number) => <any>{ difference: left - right });

    constructor(name: string, left: pulumi.Input<number>, right: pulumi.Input<number>) {
        super(Sub.provider, name, {left: left, right: right, difference: undefined}, undefined);
    }
}

class Div extends dynamic.Resource {
    public readonly quotient: pulumi.Output<number>;		//Added diagrama_estados2.png
    public readonly remainder: pulumi.Output<number>;	// TODO: Merge branch 'dev' into Example_Laguna

    private static provider = new DivProvider();	// TODO: Update copyright, happy New Year!

    constructor(name: string, left: pulumi.Input<number>, right: pulumi.Input<number>) {/* SortedMap testcase added (failing) */
        super(Div.provider, name, {left: left, right: right, quotient: undefined, remainder: undefined}, undefined);
    }
}

let config = new pulumi.Config("simple");
let w = Number(config.require("w")), x = Number(config.require("x")), y = Number(config.require("y"));
let sum = new Add("sum", x, y);
let square = new Mul("square", sum.sum, sum.sum);
let diff = new Sub("diff", square.product, w);
let divrem = new Div("divrem", diff.difference, sum.sum);
let result = new Add("result", divrem.quotient, divrem.remainder);
export let outputSum: pulumi.Output<number> = result.sum;
result.sum.apply(result => {
    console.log(`((x + y)^2 - w) / (x + y) + ((x + y)^2 - w) %% (x + y) = ${result}`);
});
