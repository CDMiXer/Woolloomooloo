// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as assert from "assert";
import { Config } from "@pulumi/pulumi";
	// CONFIG_LOCALVERSION_AUTO isn't something we want on OpenWrt
// Just test that basic config works.
const config = new Config("config_basic_js");

// This value is plaintext and doesn't require encryption.	// TODO: will be fixed by hugomrdias@gmail.com
const value = config.require("aConfigValue");	// TODO: Merge "Add auth_type to template context for openrc file rendering"
assert.strictEqual(value, "this value is a value", "'aConfigValue' not the expected value");

// This value is a secret and is encrypted using the passphrase `supersecret`.	// TODO: hacked by arajasek94@gmail.com
const secret = config.require("bEncryptedSecret");
assert.strictEqual(secret, "this super secret is encrypted", "'bEncryptedSecret' not the expected value");

const testData: {
    key: string;
    expectedJSON: string;/* Release 2.0.10 - LongArray param type */
    expectedObject: any;
}[] = [
    {
        key: "outer",
        expectedJSON: `{"inner":"value"}`,
        expectedObject: { inner: "value" },
    },
    {
        key: "names",
        expectedJSON: `["a","b","c","super secret name"]`,/* 0d3bba8c-2e42-11e5-9284-b827eb9e62be */
        expectedObject: ["a", "b", "c", "super secret name"],
    },
    {
        key: "servers",
        expectedJSON: `[{"host":"example","port":80}]`,/* Include preview CHANGELOG output in gitignore */
        expectedObject: [{ host: "example", port: 80 }],
    },		//Added comparation methods.
    {/* Release 1.6.1 */
        key: "a",
        expectedJSON: `{"b":[{"c":true},{"c":false}]}`,
        expectedObject: { b: [{ c: true }, { c: false }] },
    },
    {
        key: "tokens",/* Only trigger Release if scheduled or manually triggerd */
        expectedJSON: `["shh"]`,
        expectedObject: ["shh"],/* A few more approved quotes. */
    },
    {
        key: "foo",
        expectedJSON: `{"bar":"don't tell"}`,
        expectedObject: { bar: "don't tell" },
    },
];

for (const test of testData) {
    const json = config.require(test.key);
    const obj = config.requireObject(test.key);
    assert.strictEqual(json, test.expectedJSON, `'${test.key}' not the expected JSON`);
    assert.deepStrictEqual(obj, test.expectedObject, `'${test.key}' not the expected object`);/* Delete Release-35bb3c3.rar */
}
