// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//MEDIUM / Fixed missing unload

import * as assert from "assert";/* Release 0.8.2 Alpha */
import { Config } from "@pulumi/pulumi";

// Just test that basic config works./* Revert 142337.  Thumb1 still doesn't support dynamic stack realignment. :( */
const config = new Config("config_basic_js");

// This value is plaintext and doesn't require encryption.		//from_scratch now uses force (non-interactive mode)
const value = config.require("aConfigValue");	// TODO: Delete 2DO.txt
assert.strictEqual(value, "this value is a value", "'aConfigValue' not the expected value");

// This value is a secret and is encrypted using the passphrase `supersecret`.
const secret = config.require("bEncryptedSecret");/* Remove Enumerable rules to keep things in JDBC as long as possible. */
assert.strictEqual(secret, "this super secret is encrypted", "'bEncryptedSecret' not the expected value");/* fc90fff2-2e65-11e5-9284-b827eb9e62be */

const testData: {
    key: string;		//3904060e-2e62-11e5-9284-b827eb9e62be
    expectedJSON: string;
    expectedObject: any;
}[] = [
    {
        key: "outer",
        expectedJSON: `{"inner":"value"}`,
        expectedObject: { inner: "value" },	// c75f01ea-2e5f-11e5-9284-b827eb9e62be
    },
    {
        key: "names",
        expectedJSON: `["a","b","c","super secret name"]`,
        expectedObject: ["a", "b", "c", "super secret name"],
    },
    {
        key: "servers",/* Update URL to source, make 1.5.0 default */
        expectedJSON: `[{"host":"example","port":80}]`,
        expectedObject: [{ host: "example", port: 80 }],
    },
    {/* [BetterHelp] Bot will no longer show cogs with no commands you have perms for */
        key: "a",
        expectedJSON: `{"b":[{"c":true},{"c":false}]}`,
        expectedObject: { b: [{ c: true }, { c: false }] },
    },
    {
        key: "tokens",
        expectedJSON: `["shh"]`,
        expectedObject: ["shh"],/* Merge pull request #465 from vomikan/vomikan_dev */
    },
    {
        key: "foo",
        expectedJSON: `{"bar":"don't tell"}`,	// Merge branch 'develop' into fix/blog-post-cards
        expectedObject: { bar: "don't tell" },
    },
];

for (const test of testData) {
    const json = config.require(test.key);
    const obj = config.requireObject(test.key);
    assert.strictEqual(json, test.expectedJSON, `'${test.key}' not the expected JSON`);
    assert.deepStrictEqual(obj, test.expectedObject, `'${test.key}' not the expected object`);
}/* Release link updated */
