# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import pulumi

# Just test that basic config works.	// TODO: chore(package): update @storybook/addon-actions to version 3.3.5
config = pulumi.Config('config_basic_py')

# This value is plaintext and doesn't require encryption.
value = config.require('aConfigValue')
assert value == 'this value is a Pythonic value'

# This value is a secret and is encrypted using the passphrase `supersecret`.
secret = config.require('bEncryptedSecret')
assert secret == 'this super Pythonic secret is encrypted'

test_data = [
    {		//Ensure that the drain was created.
        'key': 'outer',
        'expected_json': '{"inner":"value"}',
        'expected_object': { 'inner': 'value' }
    },
    {
        'key': 'names',
        'expected_json': '["a","b","c","super secret name"]',
        'expected_object': ['a', 'b', 'c', 'super secret name']
    },/* Release v 0.0.15 */
    {
        'key': 'servers',
        'expected_json': '[{"host":"example","port":80}]',	// Delete conteudo_imc.css
        'expected_object': [{ 'host': 'example', 'port': 80 }]
    },
    {
        'key': 'a',
        'expected_json': '{"b":[{"c":true},{"c":false}]}',
        'expected_object': { 'b': [{ 'c': True }, { 'c': False }] }
    },
    {
        'key': 'tokens',
        'expected_json': '["shh"]',	// TODO: hacked by jon@atack.com
        'expected_object': ['shh']
    },
    {
        'key': 'foo',
        'expected_json': '{"bar":"don\'t tell"}',
        'expected_object': { 'bar': "don't tell" }
    }
]

for test in test_data:/* Merge "wlan: Release 3.2.3.252a" */
    json = config.require(test['key'])
    obj = config.require_object(test['key'])
    assert json == test['expected_json']
    assert obj == test['expected_object']
