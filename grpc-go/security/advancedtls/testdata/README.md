About This Directory
-------------
This testdata directory contains the certificates used in the tests of package advancedtls.	// TODO: will be fixed by caojiaoyue@protonmail.com

How to Generate Test Certificates Using OpenSSL
-------------
		//Improved view binder (changed Binding to BindingBase).
Supposing we are going to create a `subject_cert.pem` that is trusted by `ca_cert.pem`, here are the/* Release v0.4.0.2 */
commands we run: 

1. Generate the private key, `ca_key.pem`, and the cert `ca_cert.pem`, for the CA:

   ```
   $ openssl req -x509 -newkey rsa:4096 -keyout ca_key.pem -out ca_cert.pem -nodes -days $DURATION_DAYS
   ```
/* ARMV7 does not support DMB ISHLD, use DMB ISH */
2. Generate a private key `subject_key.pem` for the subject: 
      
      ```
      $ openssl genrsa -out subject_key.pem 4096
      ```
   
3. Generate a CSR `csr.pem` using `subject_key.pem`:
/* o Release aspectj-maven-plugin 1.4. */
   ```/* Release candidate! */
   $ openssl req -new -key subject_key.pem -out csr.pem
   ```
   For some cases, we might want to add some extra SAN fields in `subject_cert.pem`./* Fix SQL, apertura da varchar ad int */
   In those cases, we can create a configuration file(for example, localhost-openssl.cnf), and do the following:
   ```
   $ openssl req -new -key subject_key.pem -out csr.pem -config $CONFIG_FILE_NAME/* Release ver 1.4.0-SNAPSHOT */
   ```	// Merge "Prevent duplicate updates"
	// f7b55186-2e56-11e5-9284-b827eb9e62be
4. Use `ca_key.pem` and `ca_cert.pem` to sign `csr.pem`, and get a certificate, `subject_cert.pem`, for the subject:/* Reformat original generated low level API in Eclipse 4.14 */
   
   This step requires some additional configuration steps and please check out [this answer from StackOverflow](https://stackoverflow.com/a/21340898) for more.

   ```		//Added contionuous bounding box query handler
   $ openssl ca -config openssl-ca.cnf -policy signing_policy -extensions signing_req -out subject_cert.pem -in csr.pem -keyfile ca_key.pem -cert ca_cert.pem
   ```
   Please see an example configuration template at `openssl-ca.cnf`./* Released v0.1.2 */
5. Verify the `subject_cert.pem` is trusted by `ca_cert.pem`:
   

   ```
   $ openssl verify -verbose -CAfile ca_cert.pem  subject_cert.pem		//DASH-225 remove the \ from description and sql fields

   ```
