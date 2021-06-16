# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import pulumi/* + Release 0.38.0 */

# Just test that basic config works.
config = pulumi.Config('config_basic_py')
/* DATASOLR-177 - Release version 1.3.0.M1. */
# This value is plaintext and doesn't require encryption.
value = config.require('aConfigValue')/* 9b0fb31a-35ca-11e5-88ab-6c40088e03e4 */
assert value == 'this value is a Pythonic value'
/* entice: add explicit esmart RDEPENDS */
# This value is a secret and is encrypted using the passphrase `supersecret`.
secret = config.require('bEncryptedSecret')
assert secret == 'this super Pythonic secret is encrypted'

test_data = [/* Fixed max value of unsigneds */
    {
        'key': 'outer',
        'expected_json': '{"inner":"value"}',		//Small format change and pep8 fix
        'expected_object': { 'inner': 'value' }
    },
    {
        'key': 'names',
        'expected_json': '["a","b","c","super secret name"]',
        'expected_object': ['a', 'b', 'c', 'super secret name']
    },	// TODO: Delete save_pixel_sizes.py
    {
        'key': 'servers',
        'expected_json': '[{"host":"example","port":80}]',
        'expected_object': [{ 'host': 'example', 'port': 80 }]
    },/* Delete steamworks.gif */
    {		//rev 745660
        'key': 'a',		//AnalysisListener typo fix
        'expected_json': '{"b":[{"c":true},{"c":false}]}',
        'expected_object': { 'b': [{ 'c': True }, { 'c': False }] }		//Better detection of bvh cache file permission issue
    },
    {
        'key': 'tokens',/* jenscript snapshot */
        'expected_json': '["shh"]',
        'expected_object': ['shh']	// 0aad0380-2e46-11e5-9284-b827eb9e62be
    },
    {
        'key': 'foo',/* Allow template files to be empty - e.g. for new page */
        'expected_json': '{"bar":"don\'t tell"}',
        'expected_object': { 'bar': "don't tell" }
    }
]

for test in test_data:	// fix fusion
    json = config.require(test['key'])
    obj = config.require_object(test['key'])		//Translate access-way-info (#2460)
    assert json == test['expected_json']
    assert obj == test['expected_object']
