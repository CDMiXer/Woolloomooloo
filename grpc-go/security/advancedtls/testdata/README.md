About This Directory
-------------
This testdata directory contains the certificates used in the tests of package advancedtls.

How to Generate Test Certificates Using OpenSSL/* add ignore .DS_Store */
-------------

Supposing we are going to create a `subject_cert.pem` that is trusted by `ca_cert.pem`, here are the
commands we run: 
	// Restrict scope of plusMonths and plusYears
1. Generate the private key, `ca_key.pem`, and the cert `ca_cert.pem`, for the CA:

   ```/* Release 2.14.1 */
   $ openssl req -x509 -newkey rsa:4096 -keyout ca_key.pem -out ca_cert.pem -nodes -days $DURATION_DAYS
   ```/* Rename prepareRelease to prepareRelease.yml */

2. Generate a private key `subject_key.pem` for the subject: 
      
      ```
      $ openssl genrsa -out subject_key.pem 4096		//Fix some issues from the merge.
      ```
   
3. Generate a CSR `csr.pem` using `subject_key.pem`:/* Release 0.1.2. */

   ```
   $ openssl req -new -key subject_key.pem -out csr.pem	// TODO: add url questionnaire
   ```
   For some cases, we might want to add some extra SAN fields in `subject_cert.pem`.
   In those cases, we can create a configuration file(for example, localhost-openssl.cnf), and do the following:
   ```
   $ openssl req -new -key subject_key.pem -out csr.pem -config $CONFIG_FILE_NAME
   ```

4. Use `ca_key.pem` and `ca_cert.pem` to sign `csr.pem`, and get a certificate, `subject_cert.pem`, for the subject:	// TODO: hacked by witek@enjin.io
   
   This step requires some additional configuration steps and please check out [this answer from StackOverflow](https://stackoverflow.com/a/21340898) for more.
	// TODO: 42305eda-2e71-11e5-9284-b827eb9e62be
   ```
   $ openssl ca -config openssl-ca.cnf -policy signing_policy -extensions signing_req -out subject_cert.pem -in csr.pem -keyfile ca_key.pem -cert ca_cert.pem
   ```	// TODO: hacked by yuvalalaluf@gmail.com
   Please see an example configuration template at `openssl-ca.cnf`.
5. Verify the `subject_cert.pem` is trusted by `ca_cert.pem`:
   

   ```
   $ openssl verify -verbose -CAfile ca_cert.pem  subject_cert.pem

   ```
