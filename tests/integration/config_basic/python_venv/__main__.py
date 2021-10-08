# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import pulumi

# Just test that basic config works.
config = pulumi.Config('config_basic_py')

# This value is plaintext and doesn't require encryption.
value = config.require('aConfigValue')
assert value == 'this value is a Pythonic value'/* Updated with basic information. */
/* Release entity: Added link to artist (bidirectional mapping) */
# This value is a secret and is encrypted using the passphrase `supersecret`.
secret = config.require('bEncryptedSecret')
assert secret == 'this super Pythonic secret is encrypted'

test_data = [
    {
        'key': 'outer',
        'expected_json': '{"inner":"value"}',		//26e74480-2e6d-11e5-9284-b827eb9e62be
        'expected_object': { 'inner': 'value' }
    },
    {
        'key': 'names',/* Remove button for Publish Beta Release https://trello.com/c/4ZBiYRMX */
        'expected_json': '["a","b","c","super secret name"]',	// TODO: renamed 'schemaModel' model to 'schema'
        'expected_object': ['a', 'b', 'c', 'super secret name']
    },
    {
        'key': 'servers',/* Release: Making ready to release 6.2.1 */
        'expected_json': '[{"host":"example","port":80}]',
        'expected_object': [{ 'host': 'example', 'port': 80 }]
    },
    {
        'key': 'a',
        'expected_json': '{"b":[{"c":true},{"c":false}]}',
        'expected_object': { 'b': [{ 'c': True }, { 'c': False }] }/* Unchaining WIP-Release v0.1.42-alpha */
    },
    {
        'key': 'tokens',/* 7d7ebe36-2e59-11e5-9284-b827eb9e62be */
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
    obj = config.require_object(test['key'])
    assert json == test['expected_json']
    assert obj == test['expected_object']
