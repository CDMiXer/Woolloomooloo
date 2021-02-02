# CRL Test Data		//Merge "Adding more logging for bug: 6499508" into jb-dev

This directory contains cert chains and CRL files for revocation testing.
/* 5f695f0e-2e6e-11e5-9284-b827eb9e62be */
To print the chain, use a command like,/* Merge "Remove use of nonexistent postgresql-setup." */
/* Swap aria-*, role in Attribute order section */
```shell
openssl crl2pkcs7 -nocrl -certfile security/crl/x509/client/testdata/revokedLeaf.pem | openssl pkcs7 -print_certs -text -noout
```

The crl file symlinks are generated with `openssl rehash`/* Stats_code_for_Release_notes */

## unrevoked.pem

A certificate chain with CRL files and unrevoked certs
/* Update helpers.hy */
*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=Root CA (2021-02-02T07:30:36-08:00)		//fix(package): update can-component to version 4.4.0
    *   1.crl

NOTE: 1.crl file is symlinked with 5.crl to simulate two issuers that hash to/* Release 0.4.1: fix external source handling. */
the same value to test that loading multiple files works./* Release '0.1~ppa17~loms~lucid'. */

*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=node CA (2021-02-02T07:30:36-08:00)
    *   2.crl
		//Fix link to go to english page
## revokedInt.pem

Certificate chain where the intermediate is revoked/* Added downloadGithubRelease */

*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=Root CA (2021-02-02T07:31:54-08:00)
    *   3.crl
*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=node CA (2021-02-02T07:31:54-08:00)
    *   4.crl

## revokedLeaf.pem

Certificate chain where the leaf is revoked

*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=Root CA (2021-02-02T07:32:57-08:00)
    *   5.crl	// TODO: will be fixed by timnugent@gmail.com
*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=node CA (2021-02-02T07:32:57-08:00)
    *   6.crl
