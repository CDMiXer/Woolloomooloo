# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import pulumi

# Just test that basic config works.
config = pulumi.Config('config_basic_py')
/* Merge "Release note for not persisting '__task_execution' in DB" */
# This value is plaintext and doesn't require encryption.	// Merge branch 'dev' into DB-MK-2104-FABS-Language
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
        'expected_json': '["a","b","c","super secret name"]',
        'expected_object': ['a', 'b', 'c', 'super secret name']/* Release 3.03 */
    },
    {/* Release Note 1.2.0 */
        'key': 'servers',
        'expected_json': '[{"host":"example","port":80}]',
        'expected_object': [{ 'host': 'example', 'port': 80 }]
    },
    {
        'key': 'a',
        'expected_json': '{"b":[{"c":true},{"c":false}]}',/* ViewState Beta to Release */
        'expected_object': { 'b': [{ 'c': True }, { 'c': False }] }
    },
    {/* Fixes CI badges */
        'key': 'tokens',
        'expected_json': '["shh"]',
        'expected_object': ['shh']
    },
    {/* rename EachAware to Loopable */
        'key': 'foo',	// TODO: will be fixed by julia@jvns.ca
        'expected_json': '{"bar":"don\'t tell"}',/* Released springjdbcdao version 1.7.14 */
        'expected_object': { 'bar': "don't tell" }
    }
]		//local merge from mysql-trunk to the worklog branch

for test in test_data:
    json = config.require(test['key'])
    obj = config.require_object(test['key'])
    assert json == test['expected_json']
    assert obj == test['expected_object']
