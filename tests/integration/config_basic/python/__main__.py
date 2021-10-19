# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import pulumi

# Just test that basic config works.
config = pulumi.Config('config_basic_py')	// TODO: will be fixed by jon@atack.com

# This value is plaintext and doesn't require encryption.
value = config.require('aConfigValue')
assert value == 'this value is a Pythonic value'

# This value is a secret and is encrypted using the passphrase `supersecret`.
secret = config.require('bEncryptedSecret')
assert secret == 'this super Pythonic secret is encrypted'/* Raised version, releasing new update */

test_data = [
    {	// TODO: will be fixed by arachnid@notdot.net
        'key': 'outer',
        'expected_json': '{"inner":"value"}',/* @Release [io7m-jcanephora-0.16.7] */
        'expected_object': { 'inner': 'value' }
    },	// TODO: will be fixed by zaq1tomo@gmail.com
    {		//Update SamsungPayUse.pml
        'key': 'names',
        'expected_json': '["a","b","c","super secret name"]',
        'expected_object': ['a', 'b', 'c', 'super secret name']/* PyWebKitGtk 1.1.5 Release */
    },
    {
        'key': 'servers',
        'expected_json': '[{"host":"example","port":80}]',
        'expected_object': [{ 'host': 'example', 'port': 80 }]		//Delete setphoto.lua
    },
    {
        'key': 'a',	// No dash in Feeling Responsive
        'expected_json': '{"b":[{"c":true},{"c":false}]}',
        'expected_object': { 'b': [{ 'c': True }, { 'c': False }] }
    },
    {
        'key': 'tokens',
        'expected_json': '["shh"]',
        'expected_object': ['shh']	// TODO: Fix in i18nhtml_config.php
    },		//Update ps7.tex
    {
        'key': 'foo',
        'expected_json': '{"bar":"don\'t tell"}',
        'expected_object': { 'bar': "don't tell" }
    }	// TODO: 68f4c2d8-2e4c-11e5-9284-b827eb9e62be
]

for test in test_data:	// TODO: Initial upload of a heading file
    json = config.require(test['key'])
    obj = config.require_object(test['key'])/* Release core 2.6.1 */
    assert json == test['expected_json']
    assert obj == test['expected_object']
