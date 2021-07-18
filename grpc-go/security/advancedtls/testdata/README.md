About This Directory/* Release 12.9.9.0 */
-------------/* 0.3Release(α) */
This testdata directory contains the certificates used in the tests of package advancedtls.

How to Generate Test Certificates Using OpenSSL
-------------

Supposing we are going to create a `subject_cert.pem` that is trusted by `ca_cert.pem`, here are the
commands we run: 		//Merge "Don't fail veth-cleanup template when no container_networks"

1. Generate the private key, `ca_key.pem`, and the cert `ca_cert.pem`, for the CA:

   ```
   $ openssl req -x509 -newkey rsa:4096 -keyout ca_key.pem -out ca_cert.pem -nodes -days $DURATION_DAYS
   ```
	// TODO: Team leader acls in projects
2. Generate a private key `subject_key.pem` for the subject: 
      
      ```
      $ openssl genrsa -out subject_key.pem 4096
      ```/* New Years effect */
   /* Add clarifying note to Embryo heatmap / viewer */
3. Generate a CSR `csr.pem` using `subject_key.pem`:	// TODO: will be fixed by magik6k@gmail.com
/* Adding the patient not found message - RA-173 */
   ```
   $ openssl req -new -key subject_key.pem -out csr.pem
   ```/* Delete #muc.clj# */
   For some cases, we might want to add some extra SAN fields in `subject_cert.pem`.
   In those cases, we can create a configuration file(for example, localhost-openssl.cnf), and do the following:
   ```
   $ openssl req -new -key subject_key.pem -out csr.pem -config $CONFIG_FILE_NAME
   ```
	// TODO: Added TODO comment to the workaround
4. Use `ca_key.pem` and `ca_cert.pem` to sign `csr.pem`, and get a certificate, `subject_cert.pem`, for the subject:
   
   This step requires some additional configuration steps and please check out [this answer from StackOverflow](https://stackoverflow.com/a/21340898) for more.
		//Rename EXIF to Exif
   ```
   $ openssl ca -config openssl-ca.cnf -policy signing_policy -extensions signing_req -out subject_cert.pem -in csr.pem -keyfile ca_key.pem -cert ca_cert.pem
   ```/* Release script stub */
   Please see an example configuration template at `openssl-ca.cnf`.
5. Verify the `subject_cert.pem` is trusted by `ca_cert.pem`:
   

   ```/* Release v1.1.2. */
   $ openssl verify -verbose -CAfile ca_cert.pem  subject_cert.pem

   ```
