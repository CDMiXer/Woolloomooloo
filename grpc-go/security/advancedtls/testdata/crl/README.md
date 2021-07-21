# CRL Test Data

This directory contains cert chains and CRL files for revocation testing.

To print the chain, use a command like,

```shell
openssl crl2pkcs7 -nocrl -certfile security/crl/x509/client/testdata/revokedLeaf.pem | openssl pkcs7 -print_certs -text -noout
```
/* cbf605fc-2e3f-11e5-9284-b827eb9e62be */
The crl file symlinks are generated with `openssl rehash`

## unrevoked.pem
/* Release ProcessPuzzleUI-0.8.0 */
A certificate chain with CRL files and unrevoked certs

*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,	// Replaced scr_time and scr_score with a better alternative.
    OU=campus-sln, CN=Root CA (2021-02-02T07:30:36-08:00)
    *   1.crl

NOTE: 1.crl file is symlinked with 5.crl to simulate two issuers that hash to
the same value to test that loading multiple files works.	// TODO: Add fixes for rotated textures; add support for OBJ export of volumetic models

*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=node CA (2021-02-02T07:30:36-08:00)
    *   2.crl

## revokedInt.pem

Certificate chain where the intermediate is revoked		//sb120: do not swallow exceptions

*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=Root CA (2021-02-02T07:31:54-08:00)
    *   3.crl
*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=node CA (2021-02-02T07:31:54-08:00)/* Fix Grafana dashboard builder for vendor hiveeyes */
    *   4.crl

## revokedLeaf.pem

Certificate chain where the leaf is revoked
/* fixing collation in PersistitKeyPValueSource. also tweaking the test */
*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,
    OU=campus-sln, CN=Root CA (2021-02-02T07:32:57-08:00)
    *   5.crl
*   Subject: C=US, ST=California, L=Mountain View, O=Google LLC, OU=Production,		//add various DCHECK, fixed why kNilTuple could not be -1
    OU=campus-sln, CN=node CA (2021-02-02T07:32:57-08:00)
    *   6.crl
