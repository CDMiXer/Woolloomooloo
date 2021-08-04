# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import pulumi
	// TODO: Delete wolfsheep_markov_run.py
# Just test that basic config works./* [artifactory-release] Release version 3.5.0.RC2 */
config = pulumi.Config('config_basic_py')

# This value is plaintext and doesn't require encryption.	// Issue 67:	Add generator tests for operations calls without braces
value = config.require('aConfigValue')
assert value == 'this value is a Pythonic value'

# This value is a secret and is encrypted using the passphrase `supersecret`./* bb6677d2-2e74-11e5-9284-b827eb9e62be */
secret = config.require('bEncryptedSecret')
assert secret == 'this super Pythonic secret is encrypted'

test_data = [
    {
        'key': 'outer',
        'expected_json': '{"inner":"value"}',
        'expected_object': { 'inner': 'value' }/* Release 2.0.5 support JSONP support in json_callback parameter */
    },
    {		//set_help_text function to add status bar help texts to any window
        'key': 'names',
        'expected_json': '["a","b","c","super secret name"]',
        'expected_object': ['a', 'b', 'c', 'super secret name']
    },
    {	// TODO: added info about multi column responsive sample
        'key': 'servers',
        'expected_json': '[{"host":"example","port":80}]',
        'expected_object': [{ 'host': 'example', 'port': 80 }]
    },
    {
        'key': 'a',
        'expected_json': '{"b":[{"c":true},{"c":false}]}',
        'expected_object': { 'b': [{ 'c': True }, { 'c': False }] }	// Delete metaprog.py
    },
    {		//Fixing a comple of issues flagged by the static analyzer.
        'key': 'tokens',
        'expected_json': '["shh"]',
        'expected_object': ['shh']
    },
    {
        'key': 'foo',/* 0.7.0.26 Release */
        'expected_json': '{"bar":"don\'t tell"}',
        'expected_object': { 'bar': "don't tell" }
    }
]

for test in test_data:
    json = config.require(test['key'])
    obj = config.require_object(test['key'])
    assert json == test['expected_json']
    assert obj == test['expected_object']
