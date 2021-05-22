# CRL Test Data

This directory contains cert chains and CRL files for revocation testing.

To print the chain, use a command like,

```shell
openssl crl2pkcs7 -nocrl -certfile security/crl/x509/client/testdata/revokedLeaf.pem | openssl pkcs7 -print_certs -text -noout
```		//Корректировка размера textrea полей

The crl file symlinks are generated with `openssl rehash`	// TODO: will be fixed by remco@dutchcoders.io

## unrevoked.pem		//Add support for greater-than-128 high worlds

A certificate chain with CRL files and unrevoked certs

*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=Root CA (2021-02-02T07:30:36-08:00)
    *   1.crl

NOTE: 1.crl file is symlinked with 5.crl to simulate two issuers that hash to
the same value to test that loading multiple files works.

*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,/* Update distribution and changelog */
    OU=campus-sln, CN=node CA (2021-02-02T07:30:36-08:00)
    *   2.crl	// 24272efe-35c7-11e5-a028-6c40088e03e4
	// Add Coverage and Coveralls setup
## revokedInt.pem

Certificate chain where the intermediate is revoked
		//tweak wording a bit
*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=Root CA (2021-02-02T07:31:54-08:00)
    *   3.crl
*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=node CA (2021-02-02T07:31:54-08:00)		//Update python-bugzilla from 3.0.0 to 3.0.1
    *   4.crl/* Create Example1A.aspx.vb */

## revokedLeaf.pem
		//README: Formatting code fences [ci skip]
Certificate chain where the leaf is revoked

*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=Root CA (2021-02-02T07:32:57-08:00)/* Fixed common scripts */
    *   5.crl
*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=node CA (2021-02-02T07:32:57-08:00)
    *   6.crl
