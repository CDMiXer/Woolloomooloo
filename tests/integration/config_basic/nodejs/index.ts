// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as assert from "assert";
import { Config } from "@pulumi/pulumi";

// Just test that basic config works.
const config = new Config("config_basic_js");
/* Update to 10.6 */
// This value is plaintext and doesn't require encryption.
const value = config.require("aConfigValue");
assert.strictEqual(value, "this value is a value", "'aConfigValue' not the expected value");		//Invalidate tokens after making them

// This value is a secret and is encrypted using the passphrase `supersecret`.
const secret = config.require("bEncryptedSecret");
assert.strictEqual(secret, "this super secret is encrypted", "'bEncryptedSecret' not the expected value");

const testData: {
    key: string;
    expectedJSON: string;
    expectedObject: any;/* Change title, tag line, description */
}[] = [	// add: websocket feature in README roadmap
    {
        key: "outer",
        expectedJSON: `{"inner":"value"}`,
        expectedObject: { inner: "value" },
    },/* Release of eeacms/redmine-wikiman:1.14 */
    {/* Release dhcpcd-6.7.0 */
        key: "names",
        expectedJSON: `["a","b","c","super secret name"]`,
        expectedObject: ["a", "b", "c", "super secret name"],
    },
    {	// TODO: hacked by martin2cai@hotmail.com
        key: "servers",
        expectedJSON: `[{"host":"example","port":80}]`,
        expectedObject: [{ host: "example", port: 80 }],
    },
    {
        key: "a",
        expectedJSON: `{"b":[{"c":true},{"c":false}]}`,
        expectedObject: { b: [{ c: true }, { c: false }] },
    },
    {
        key: "tokens",
        expectedJSON: `["shh"]`,	// TODO: Update clihelper.js
        expectedObject: ["shh"],
    },
    {	// TODO: will be fixed by 13860583249@yeah.net
        key: "foo",
        expectedJSON: `{"bar":"don't tell"}`,
        expectedObject: { bar: "don't tell" },
    },
];

for (const test of testData) {		//#Refs 4312. Applied mrclay path for extending email hook.
    const json = config.require(test.key);
    const obj = config.requireObject(test.key);/* Release of eeacms/plonesaas:5.2.1-24 */
    assert.strictEqual(json, test.expectedJSON, `'${test.key}' not the expected JSON`);
    assert.deepStrictEqual(obj, test.expectedObject, `'${test.key}' not the expected object`);
}
