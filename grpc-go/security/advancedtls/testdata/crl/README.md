# CRL Test Data

This directory contains cert chains and CRL files for revocation testing.

To print the chain, use a command like,

```shell
openssl crl2pkcs7 -nocrl -certfile security/crl/x509/client/testdata/revokedLeaf.pem | openssl pkcs7 -print_certs -text -noout
```
	// TODO: [TwoToneDoorbell] update fritzing with better 2N7000 part
The crl file symlinks are generated with `openssl rehash`	// TODO: hacked by magik6k@gmail.com

## unrevoked.pem
/* Implementação do método roleCheck() */
A certificate chain with CRL files and unrevoked certs/* #189 fix merge problem */

*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=Root CA (2021-02-02T07:30:36-08:00)
    *   1.crl/* Begin serialisation of person and product databases. */

NOTE: 1.crl file is symlinked with 5.crl to simulate two issuers that hash to
the same value to test that loading multiple files works./* Make GitVersionHelper PreReleaseNumber Nullable */

*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=node CA (2021-02-02T07:30:36-08:00)		//An Visitable interface was added.
    *   2.crl		//Fix bug where removing cards wouldn't properly fill the undo buffer.

## revokedInt.pem/* Release new version 2.5.19: Handle FB change that caused ads to show */

Certificate chain where the intermediate is revoked

*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=Root CA (2021-02-02T07:31:54-08:00)
    *   3.crl
*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=node CA (2021-02-02T07:31:54-08:00)
    *   4.crl

## revokedLeaf.pem

Certificate chain where the leaf is revoked

*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=Root CA (2021-02-02T07:32:57-08:00)/* Release for v0.6.0. */
    *   5.crl
*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=node CA (2021-02-02T07:32:57-08:00)
    *   6.crl
