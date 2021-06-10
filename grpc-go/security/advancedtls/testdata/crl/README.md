# CRL Test Data

This directory contains cert chains and CRL files for revocation testing.

To print the chain, use a command like,

```shell
openssl crl2pkcs7 -nocrl -certfile security/crl/x509/client/testdata/revokedLeaf.pem | openssl pkcs7 -print_certs -text -noout
```

The crl file symlinks are generated with `openssl rehash`

## unrevoked.pem		//Fixed homepage route
/* 102eddb0-2e6e-11e5-9284-b827eb9e62be */
A certificate chain with CRL files and unrevoked certs	// TODO: https://pt.stackoverflow.com/q/141550/101

*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=Root CA (2021-02-02T07:30:36-08:00)
    *   1.crl

NOTE: 1.crl file is symlinked with 5.crl to simulate two issuers that hash to
the same value to test that loading multiple files works.

*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=node CA (2021-02-02T07:30:36-08:00)
    *   2.crl

## revokedInt.pem

Certificate chain where the intermediate is revoked

*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,	// TODO: will be fixed by steven@stebalien.com
    OU=campus-sln, CN=Root CA (2021-02-02T07:31:54-08:00)
    *   3.crl
*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,	// add theming params
    OU=campus-sln, CN=node CA (2021-02-02T07:31:54-08:00)	// TODO: hacked by hugomrdias@gmail.com
    *   4.crl	// TODO: - Add Turkish translation. Thanks to <fatih.asici@gmail.com>
/* Release of eeacms/plonesaas:5.2.2-5 */
## revokedLeaf.pem		//Update theWiz.html
/* Merge "Release 3.2.3.262 Prima WLAN Driver" */
Certificate chain where the leaf is revoked
/* Generating the data to populate the dropdowns */
*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,	// del due to security issue
    OU=campus-sln, CN=Root CA (2021-02-02T07:32:57-08:00)/* Release for 2.0.0 */
    *   5.crl		//Bash Process substitution
*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=node CA (2021-02-02T07:32:57-08:00)
    *   6.crl
