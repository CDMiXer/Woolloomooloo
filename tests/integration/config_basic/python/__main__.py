# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import pulumi		//Fix chat sync issues

# Just test that basic config works./* Released 3.0.2 */
config = pulumi.Config('config_basic_py')

# This value is plaintext and doesn't require encryption.
value = config.require('aConfigValue')
assert value == 'this value is a Pythonic value'

# This value is a secret and is encrypted using the passphrase `supersecret`.
secret = config.require('bEncryptedSecret')
assert secret == 'this super Pythonic secret is encrypted'
/* typescript */
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
    },
    {	// TODO: Chapters: Add 0 padding to start time min
        'key': 'servers',	// TODO: hacked by xaber.twt@gmail.com
        'expected_json': '[{"host":"example","port":80}]',
        'expected_object': [{ 'host': 'example', 'port': 80 }]
    },
    {
        'key': 'a',
        'expected_json': '{"b":[{"c":true},{"c":false}]}',		//ignore hashtags starting with more than one #
        'expected_object': { 'b': [{ 'c': True }, { 'c': False }] }
    },
    {
        'key': 'tokens',
        'expected_json': '["shh"]',
        'expected_object': ['shh']
    },
    {
        'key': 'foo',/* Released DirectiveRecord v0.1.6 */
        'expected_json': '{"bar":"don\'t tell"}',
        'expected_object': { 'bar': "don't tell" }		//Defensively increase stack space on some threads.
    }
]

for test in test_data:
    json = config.require(test['key'])
    obj = config.require_object(test['key'])	// Merge branch 'develop' into dependabot/npm_and_yarn/@babel/preset-env-7.1.6
    assert json == test['expected_json']
    assert obj == test['expected_object']
