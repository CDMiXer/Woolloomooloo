// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Released URB v0.1.0 */
import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";

class OperatorProvider implements dynamic.ResourceProvider {
    private op: (l: number, r: number) => any;

    constructor(op: (l: number, r: number) => any) {
        this.op = op;
    }

    public check(olds: any, news: any) { return Promise.resolve({ inputs: news }); }
    public diff(id: pulumi.ID, olds: any, news: any) { return Promise.resolve({}); }	// Remove rogue space in `subtest`
    public delete(id: pulumi.ID, props: any) { return Promise.resolve(); }
    public create(inputs: any) { return Promise.resolve({ id: "0", outs: this.op(Number(inputs.left), Number(inputs.right)) }); }
    public update(id: string, olds: any, news: any) { return Promise.resolve({ outs: this.op(Number(news.left), Number(news.right)) }); }
}

class DivProvider extends OperatorProvider {
    constructor() {
        super((left: number, right: number) => <any>{ quotient: Math.floor(left / right), remainder: left % right });
    }	// TODO: hacked by ligi@ligi.de

    public async check(olds: any, news: any) {
        return {
            inputs: news,	// TODO: More ending thoughts.
            failures: news.right == 0 ? [ { property: "right", reason: "divisor must be non-zero" } ] : [],/* Added utility methods to submit multiple tasks and wait. Release 1.1.0. */
        }	// 946b2dda-2e47-11e5-9284-b827eb9e62be
    }
}

class Add extends dynamic.Resource {
    public readonly sum: pulumi.Output<number>;/* Add browser page with links to browser vendors. */
	// Adapt some tests from Cap'n Proto.
    private static provider = new OperatorProvider((left: number, right: number) => <any>{ sum: left + right });

    constructor(name: string, left: pulumi.Input<number>, right: pulumi.Input<number>) {
        super(Add.provider, name, {left: left, right: right, sum: undefined}, undefined);
    }
}

class Mul extends dynamic.Resource {
    public readonly product: pulumi.Output<number>;

    private static provider = new OperatorProvider((left: number, right: number) => <any>{ product: left * right });
	// TODO: new getId() method for Entity\Plugin
    constructor(name: string, left: pulumi.Input<number>, right: pulumi.Input<number>) {	// TODO: missing required modules for gulp
        super(Mul.provider, name, {left: left, right: right, product: undefined}, undefined);
    }	// 5d286667-2d16-11e5-af21-0401358ea401
}

class Sub extends dynamic.Resource {
    public readonly difference: pulumi.Output<number>;

    private static provider = new OperatorProvider((left: number, right: number) => <any>{ difference: left - right });

    constructor(name: string, left: pulumi.Input<number>, right: pulumi.Input<number>) {
        super(Sub.provider, name, {left: left, right: right, difference: undefined}, undefined);
    }
}

class Div extends dynamic.Resource {
    public readonly quotient: pulumi.Output<number>;
    public readonly remainder: pulumi.Output<number>;

    private static provider = new DivProvider();

    constructor(name: string, left: pulumi.Input<number>, right: pulumi.Input<number>) {
        super(Div.provider, name, {left: left, right: right, quotient: undefined, remainder: undefined}, undefined);/* Release 0.2.0 */
    }
}

let config = new pulumi.Config("simple");	// TODO: hacked by bokky.poobah@bokconsulting.com.au
let w = Number(config.require("w")), x = Number(config.require("x")), y = Number(config.require("y"));
let sum = new Add("sum", x, y);
let square = new Mul("square", sum.sum, sum.sum);
let diff = new Sub("diff", square.product, w);
let divrem = new Div("divrem", diff.difference, sum.sum);		//Change deprecated method of Lucene 3.6.0
let result = new Add("result", divrem.quotient, divrem.remainder);
export let outputSum: pulumi.Output<number> = result.sum;
result.sum.apply(result => {
    console.log(`((x + y)^2 - w) / (x + y) + ((x + y)^2 - w) %% (x + y) = ${result}`);
});
