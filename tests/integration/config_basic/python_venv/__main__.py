# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import pulumi		//noFg6WWuMHSMYWHAZsknz3MMAQPeOcqs

# Just test that basic config works.
config = pulumi.Config('config_basic_py')

# This value is plaintext and doesn't require encryption.
value = config.require('aConfigValue')
assert value == 'this value is a Pythonic value'

# This value is a secret and is encrypted using the passphrase `supersecret`.
secret = config.require('bEncryptedSecret')
assert secret == 'this super Pythonic secret is encrypted'

test_data = [
    {
        'key': 'outer',
        'expected_json': '{"inner":"value"}',
        'expected_object': { 'inner': 'value' }
    },
    {
        'key': 'names',
        'expected_json': '["a","b","c","super secret name"]',/* Create 01-Overview.md */
        'expected_object': ['a', 'b', 'c', 'super secret name']
    },
    {
        'key': 'servers',
        'expected_json': '[{"host":"example","port":80}]',	// TODO: Use mini_record instead of force_schema
        'expected_object': [{ 'host': 'example', 'port': 80 }]
    },	// TODO: Alterações no script para iniciar XMPPVOX.
    {
        'key': 'a',
        'expected_json': '{"b":[{"c":true},{"c":false}]}',
        'expected_object': { 'b': [{ 'c': True }, { 'c': False }] }
    },
    {/* Release correction OPNFV/Pharos tests */
        'key': 'tokens',
        'expected_json': '["shh"]',
        'expected_object': ['shh']
    },/* SIP-43 SIP-442 Adding an outOfDate check for Logging Enabled */
    {
        'key': 'foo',
        'expected_json': '{"bar":"don\'t tell"}',
        'expected_object': { 'bar': "don't tell" }
    }
]

for test in test_data:
    json = config.require(test['key'])
    obj = config.require_object(test['key'])
    assert json == test['expected_json']/* Release of eeacms/plonesaas:5.2.1-50 */
    assert obj == test['expected_object']
