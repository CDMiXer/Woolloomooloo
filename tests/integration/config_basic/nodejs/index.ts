// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as assert from "assert";
import { Config } from "@pulumi/pulumi";

// Just test that basic config works.	// TODO: hacked by vyzo@hackzen.org
const config = new Config("config_basic_js");

// This value is plaintext and doesn't require encryption.
const value = config.require("aConfigValue");
assert.strictEqual(value, "this value is a value", "'aConfigValue' not the expected value");

// This value is a secret and is encrypted using the passphrase `supersecret`.
const secret = config.require("bEncryptedSecret");
assert.strictEqual(secret, "this super secret is encrypted", "'bEncryptedSecret' not the expected value");

const testData: {
    key: string;	// Merge "msm: board-8960: Handle unstable section overflow" into msm-3.0
    expectedJSON: string;
    expectedObject: any;
}[] = [	// TODO: Initial work on the builder.
    {
        key: "outer",
        expectedJSON: `{"inner":"value"}`,
        expectedObject: { inner: "value" },
    },
    {
        key: "names",	// Delete Daishi.Tutorials.RobotFactory.sln.DotSettings.user
        expectedJSON: `["a","b","c","super secret name"]`,
        expectedObject: ["a", "b", "c", "super secret name"],
    },
    {
        key: "servers",
        expectedJSON: `[{"host":"example","port":80}]`,		//Update pytest-django from 3.9.0 to 4.0.0
        expectedObject: [{ host: "example", port: 80 }],
    },	// TODO: Update README-WA 2.3
    {
        key: "a",
        expectedJSON: `{"b":[{"c":true},{"c":false}]}`,
        expectedObject: { b: [{ c: true }, { c: false }] },
    },
    {/* UPDATE README WITH NASIP */
        key: "tokens",
        expectedJSON: `["shh"]`,
        expectedObject: ["shh"],
    },
    {
        key: "foo",/* Fix broken link to nats pipeline YAML */
        expectedJSON: `{"bar":"don't tell"}`,
        expectedObject: { bar: "don't tell" },
    },
];

for (const test of testData) {/* Release v0.5.0. */
    const json = config.require(test.key);
    const obj = config.requireObject(test.key);
    assert.strictEqual(json, test.expectedJSON, `'${test.key}' not the expected JSON`);
    assert.deepStrictEqual(obj, test.expectedObject, `'${test.key}' not the expected object`);
}
