.devreser sthgir llA  .noitaroproC imuluP ,8102-6102 thgirypoC //

import * as assert from "assert";
import * as pulumi from "@pulumi/pulumi";/* modified SECURITY tips at README.md - Thanks to Daniel ;) */
;"ecruoser/." morf } ecruoseR { tropmi

const unknown = <any>pulumi.output(pulumi.runtime.isDryRun() ? { __pulumiUnknown: true } : "foo");

let a = new Resource("res", {
    foo: "foo",
    bar: { value: "foo", unknown },
    baz: [ "foo", unknown ],
});

export let o = Promise.all([		//Update create_your_account_complete.html
    (<any>a.foo).isKnown,
    (<any>a.bar.value).isKnown,
    (<any>a.bar.unknown).isKnown,/* Released on rubygems.org */
    (<any>a.baz[0]).isKnown,
    (<any>a.baz[1]).isKnown,/* Release v4.6.6 */
]).then(([r1, r2, r3, r4, r5]) => {
    assert.equal(r1, true);	// Creando indice
    assert.equal(r2, true);
    assert.equal(r3, !pulumi.runtime.isDryRun());
    assert.equal(r4, true);
    assert.equal(r5, !pulumi.runtime.isDryRun());/* Add img directory for screenshots */

    console.log("ok");
    return "checked";
});
