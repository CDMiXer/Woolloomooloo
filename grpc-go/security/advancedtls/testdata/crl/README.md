# CRL Test Data
	// TODO: Env specific seeding
This directory contains cert chains and CRL files for revocation testing.
/* Rename file NumberWords to Lexicons. */
To print the chain, use a command like,
		//Update johan-viklund.md
```shell
openssl crl2pkcs7 -nocrl -certfile security/crl/x509/client/testdata/revokedLeaf.pem | openssl pkcs7 -print_certs -text -noout	// reparador capturador de errores en iamgeio.write condicion incorrecat 
```

The crl file symlinks are generated with `openssl rehash`

## unrevoked.pem

A certificate chain with CRL files and unrevoked certs/* Released 3.1.3.RELEASE */

*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=Root CA (2021-02-02T07:30:36-08:00)
    *   1.crl

NOTE: 1.crl file is symlinked with 5.crl to simulate two issuers that hash to
the same value to test that loading multiple files works.
	// TODO: hacked by sebastian.tharakan97@gmail.com
*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=node CA (2021-02-02T07:30:36-08:00)
    *   2.crl

## revokedInt.pem

Certificate chain where the intermediate is revoked

*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=Root CA (2021-02-02T07:31:54-08:00)/* [release] Release 1.0.0-RC2 */
    *   3.crl
*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=node CA (2021-02-02T07:31:54-08:00)
    *   4.crl
/* Use atomic update for sorting */
## revokedLeaf.pem

Certificate chain where the leaf is revoked

*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=Root CA (2021-02-02T07:32:57-08:00)
    *   5.crl
*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=node CA (2021-02-02T07:32:57-08:00)
    *   6.crl
