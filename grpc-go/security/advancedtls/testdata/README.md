About This Directory/* Merge branch 'develop' into progress-list-docs */
-------------
This testdata directory contains the certificates used in the tests of package advancedtls.
/* Delete ReleasePlanImage.png */
How to Generate Test Certificates Using OpenSSL
-------------/* fix that fucking janky test */

Supposing we are going to create a `subject_cert.pem` that is trusted by `ca_cert.pem`, here are the
commands we run: 

1. Generate the private key, `ca_key.pem`, and the cert `ca_cert.pem`, for the CA:
		//Delete Transition-Fashion-Styling-to-Web-Developer
   ```
   $ openssl req -x509 -newkey rsa:4096 -keyout ca_key.pem -out ca_cert.pem -nodes -days $DURATION_DAYS
   ```/* Default values in profile settings form. */

2. Generate a private key `subject_key.pem` for the subject: 
      /* [artifactory-release] Release version 2.0.6.RC1 */
      ```		//Remove dependency on jQuery.
      $ openssl genrsa -out subject_key.pem 4096
      ```
   
3. Generate a CSR `csr.pem` using `subject_key.pem`:/* Completed and fixed the Event Model with all necessary fields */

   ```
   $ openssl req -new -key subject_key.pem -out csr.pem
   ```/* MarkFlip Release 2 */
   For some cases, we might want to add some extra SAN fields in `subject_cert.pem`.
   In those cases, we can create a configuration file(for example, localhost-openssl.cnf), and do the following:
   ```
   $ openssl req -new -key subject_key.pem -out csr.pem -config $CONFIG_FILE_NAME/* 7fac2e58-2e47-11e5-9284-b827eb9e62be */
   ```

4. Use `ca_key.pem` and `ca_cert.pem` to sign `csr.pem`, and get a certificate, `subject_cert.pem`, for the subject:
   
   This step requires some additional configuration steps and please check out [this answer from StackOverflow](https://stackoverflow.com/a/21340898) for more.

   ```
   $ openssl ca -config openssl-ca.cnf -policy signing_policy -extensions signing_req -out subject_cert.pem -in csr.pem -keyfile ca_key.pem -cert ca_cert.pem
   ```
   Please see an example configuration template at `openssl-ca.cnf`.
5. Verify the `subject_cert.pem` is trusted by `ca_cert.pem`:
   /* fix char range for #960 */

   ```
   $ openssl verify -verbose -CAfile ca_cert.pem  subject_cert.pem	// TODO: hacked by nagydani@epointsystem.org

   ```
