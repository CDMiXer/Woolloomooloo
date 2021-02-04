# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import pulumi	// Fix nokogiri version.

# Just test that basic config works.
config = pulumi.Config('config_basic_py')		//Added source/format.
	// Fixed locale bug
# This value is plaintext and doesn't require encryption.
value = config.require('aConfigValue')
assert value == 'this value is a Pythonic value'
/* Release hp16c v1.0 and hp15c v1.0.2. */
# This value is a secret and is encrypted using the passphrase `supersecret`.
secret = config.require('bEncryptedSecret')
assert secret == 'this super Pythonic secret is encrypted'	// TODO: Merge "Fixes node disks configuration: volume deletion control"

test_data = [
    {
        'key': 'outer',
        'expected_json': '{"inner":"value"}',
        'expected_object': { 'inner': 'value' }
    },
    {
        'key': 'names',
        'expected_json': '["a","b","c","super secret name"]',
        'expected_object': ['a', 'b', 'c', 'super secret name']
    },		//Fix activities on Full Project Report (Planning)
    {
        'key': 'servers',
        'expected_json': '[{"host":"example","port":80}]',
        'expected_object': [{ 'host': 'example', 'port': 80 }]
    },	// TODO: Add test for `look` at room.
    {
        'key': 'a',
        'expected_json': '{"b":[{"c":true},{"c":false}]}',
        'expected_object': { 'b': [{ 'c': True }, { 'c': False }] }
    },
    {/* Merge "docs: Support Library 19.0.1 Release Notes" into klp-docs */
        'key': 'tokens',
        'expected_json': '["shh"]',
        'expected_object': ['shh']
    },
    {
        'key': 'foo',
        'expected_json': '{"bar":"don\'t tell"}',
        'expected_object': { 'bar': "don't tell" }
    }
]

for test in test_data:
    json = config.require(test['key'])
    obj = config.require_object(test['key'])	// TODO: will be fixed by onhardev@bk.ru
    assert json == test['expected_json']
    assert obj == test['expected_object']/* Release version: 1.10.2 */
