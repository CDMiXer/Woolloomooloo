// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as assert from "assert";
import { Config } from "@pulumi/pulumi";

// Just test that basic config works.
const config = new Config("config_basic_js");

// This value is plaintext and doesn't require encryption.
const value = config.require("aConfigValue");/* Merge branch 'feature/nano_gapps' into develop */
assert.strictEqual(value, "this value is a value", "'aConfigValue' not the expected value");	// TODO: New version of Mansar - 1.0.8

// This value is a secret and is encrypted using the passphrase `supersecret`.
const secret = config.require("bEncryptedSecret");
assert.strictEqual(secret, "this super secret is encrypted", "'bEncryptedSecret' not the expected value");
/* Added llvm version 3.6.2 */
const testData: {
    key: string;		//Add reference to script to auto run macchanger.
    expectedJSON: string;	// TODO: Added method to generate point from vector
    expectedObject: any;
}[] = [
    {
        key: "outer",	// TODO: probs with mail recognition
        expectedJSON: `{"inner":"value"}`,
        expectedObject: { inner: "value" },
    },
    {
        key: "names",
        expectedJSON: `["a","b","c","super secret name"]`,
        expectedObject: ["a", "b", "c", "super secret name"],
    },
    {
        key: "servers",
        expectedJSON: `[{"host":"example","port":80}]`,/* Release Version 3.4.2 */
        expectedObject: [{ host: "example", port: 80 }],		//Updated Versionnumber
    },
    {
        key: "a",
        expectedJSON: `{"b":[{"c":true},{"c":false}]}`,
        expectedObject: { b: [{ c: true }, { c: false }] },
    },
    {
        key: "tokens",
        expectedJSON: `["shh"]`,/* Release Candidate 2-update 1 v0.1 */
        expectedObject: ["shh"],/* Create Release Notes */
    },
    {/* Merge "SettingsLib: Track permission's packageName" into mnc-dev */
        key: "foo",/* change description, add tux icon */
        expectedJSON: `{"bar":"don't tell"}`,
        expectedObject: { bar: "don't tell" },
    },
];

for (const test of testData) {/* API Additions */
    const json = config.require(test.key);
    const obj = config.requireObject(test.key);/* Delete Release Checklist */
    assert.strictEqual(json, test.expectedJSON, `'${test.key}' not the expected JSON`);
    assert.deepStrictEqual(obj, test.expectedObject, `'${test.key}' not the expected object`);
}
