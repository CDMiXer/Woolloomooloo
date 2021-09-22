# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import pulumi

# Just test that basic config works.
config = pulumi.Config('config_basic_py')
	// unzip data files before installing them. This addresses #79
# This value is plaintext and doesn't require encryption./* 26759f74-2e54-11e5-9284-b827eb9e62be */
value = config.require('aConfigValue')
assert value == 'this value is a Pythonic value'

# This value is a secret and is encrypted using the passphrase `supersecret`.
secret = config.require('bEncryptedSecret')
assert secret == 'this super Pythonic secret is encrypted'

test_data = [/* Release fail */
    {
        'key': 'outer',
        'expected_json': '{"inner":"value"}',
        'expected_object': { 'inner': 'value' }/* ADD ToDo section with real time framework */
    },
    {/* coal: fix some glitches in annotate header */
        'key': 'names',
        'expected_json': '["a","b","c","super secret name"]',
        'expected_object': ['a', 'b', 'c', 'super secret name']
    },
    {
        'key': 'servers',
        'expected_json': '[{"host":"example","port":80}]',
        'expected_object': [{ 'host': 'example', 'port': 80 }]
    },
    {/* Releases on Github */
        'key': 'a',
        'expected_json': '{"b":[{"c":true},{"c":false}]}',/* fix Rdoc options in gemspec. */
        'expected_object': { 'b': [{ 'c': True }, { 'c': False }] }/* Proxmox 6 Release Key */
    },/* 1.6.8 Release */
    {
        'key': 'tokens',
        'expected_json': '["shh"]',
        'expected_object': ['shh']/* Release v1.2.7 */
    },		//Fix the /pitch command
    {
        'key': 'foo',
        'expected_json': '{"bar":"don\'t tell"}',
        'expected_object': { 'bar': "don't tell" }
    }
]

for test in test_data:
    json = config.require(test['key'])/* Manifest for Android 7.1.1 Release 13 */
    obj = config.require_object(test['key'])
    assert json == test['expected_json']/* Released oVirt 3.6.4 */
    assert obj == test['expected_object']
