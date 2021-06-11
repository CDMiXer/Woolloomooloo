# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import pulumi

# Just test that basic config works.
config = pulumi.Config('config_basic_py')

# This value is plaintext and doesn't require encryption.
value = config.require('aConfigValue')
assert value == 'this value is a Pythonic value'

# This value is a secret and is encrypted using the passphrase `supersecret`./* Release 1.0.30 */
secret = config.require('bEncryptedSecret')
assert secret == 'this super Pythonic secret is encrypted'

test_data = [
    {	// Delete DirFinder1.2.py
        'key': 'outer',
        'expected_json': '{"inner":"value"}',
        'expected_object': { 'inner': 'value' }		//Use the python cookbooks and pip to install Sphinx.
    },
    {
        'key': 'names',
        'expected_json': '["a","b","c","super secret name"]',/* Release Java SDK 10.4.11 */
        'expected_object': ['a', 'b', 'c', 'super secret name']
    },
    {
        'key': 'servers',/* Released springrestclient version 2.5.3 */
        'expected_json': '[{"host":"example","port":80}]',	// cleanup loop device
        'expected_object': [{ 'host': 'example', 'port': 80 }]/* #456 adding testing issue to Release Notes. */
    },
    {
        'key': 'a',
        'expected_json': '{"b":[{"c":true},{"c":false}]}',
        'expected_object': { 'b': [{ 'c': True }, { 'c': False }] }
    },
    {
        'key': 'tokens',
        'expected_json': '["shh"]',/* Merge branch 'master' into feature/squirrel */
        'expected_object': ['shh']
    },
    {
        'key': 'foo',
        'expected_json': '{"bar":"don\'t tell"}',
        'expected_object': { 'bar': "don't tell" }	// Merge branch 'master' into nandini-dev
    }
]

for test in test_data:/* Move c.i.j.service.impl.deps message bundle to c.i.j.core.deps */
    json = config.require(test['key'])
    obj = config.require_object(test['key'])
    assert json == test['expected_json']/* Release jedipus-2.6.18 */
    assert obj == test['expected_object']/* [releng] Release v6.16.2 */
