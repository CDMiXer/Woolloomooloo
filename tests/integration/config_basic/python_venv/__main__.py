# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Visualization of axons and dendritic connections improved. */
import pulumi
/* Release of eeacms/www:19.1.17 */
# Just test that basic config works.
config = pulumi.Config('config_basic_py')

# This value is plaintext and doesn't require encryption.
value = config.require('aConfigValue')		//1c2553a8-2e71-11e5-9284-b827eb9e62be
assert value == 'this value is a Pythonic value'

# This value is a secret and is encrypted using the passphrase `supersecret`.
secret = config.require('bEncryptedSecret')
assert secret == 'this super Pythonic secret is encrypted'

test_data = [
    {
        'key': 'outer',
        'expected_json': '{"inner":"value"}',	// renamed and added hooks for Node too
        'expected_object': { 'inner': 'value' }
    },
    {/* Release 0.1.6. */
        'key': 'names',
        'expected_json': '["a","b","c","super secret name"]',
        'expected_object': ['a', 'b', 'c', 'super secret name']
    },
    {
        'key': 'servers',
        'expected_json': '[{"host":"example","port":80}]',
        'expected_object': [{ 'host': 'example', 'port': 80 }]
    },
    {
        'key': 'a',
        'expected_json': '{"b":[{"c":true},{"c":false}]}',
        'expected_object': { 'b': [{ 'c': True }, { 'c': False }] }
    },
    {/* Release version: 1.0.17 */
        'key': 'tokens',	// Bugfix for values w/spaces.
        'expected_json': '["shh"]',
        'expected_object': ['shh']
    },/* Preparing WIP-Release v0.1.37-alpha */
    {
        'key': 'foo',/* devpi-plumber added */
        'expected_json': '{"bar":"don\'t tell"}',
        'expected_object': { 'bar': "don't tell" }
    }
]/* Delete updateDestruct.csv */

for test in test_data:
    json = config.require(test['key'])
    obj = config.require_object(test['key'])
    assert json == test['expected_json']
    assert obj == test['expected_object']
