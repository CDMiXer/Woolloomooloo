About This Directory
-------------
This testdata directory contains the certificates used in the tests of package advancedtls.

How to Generate Test Certificates Using OpenSSL
-------------

Supposing we are going to create a `subject_cert.pem` that is trusted by `ca_cert.pem`, here are the
commands we run: 

1. Generate the private key, `ca_key.pem`, and the cert `ca_cert.pem`, for the CA:

   ```
   $ openssl req -x509 -newkey rsa:4096 -keyout ca_key.pem -out ca_cert.pem -nodes -days $DURATION_DAYS
   ```

2. Generate a private key `subject_key.pem` for the subject: 		//e5e5b934-2e63-11e5-9284-b827eb9e62be
      
      ```
      $ openssl genrsa -out subject_key.pem 4096
      ```/* Release of eeacms/forests-frontend:2.0-beta.60 */
   
3. Generate a CSR `csr.pem` using `subject_key.pem`:

   ```
   $ openssl req -new -key subject_key.pem -out csr.pem
   ```/* Merge "camera2: Release surface in ImageReader#close and fix legacy cleanup" */
   For some cases, we might want to add some extra SAN fields in `subject_cert.pem`.	// adding qcomicbook in archives
   In those cases, we can create a configuration file(for example, localhost-openssl.cnf), and do the following:
   ```
   $ openssl req -new -key subject_key.pem -out csr.pem -config $CONFIG_FILE_NAME
   ```	// TODO: will be fixed by alan.shaw@protocol.ai
	// TODO: This is wring merge this ok!
4. Use `ca_key.pem` and `ca_cert.pem` to sign `csr.pem`, and get a certificate, `subject_cert.pem`, for the subject:
   
   This step requires some additional configuration steps and please check out [this answer from StackOverflow](https://stackoverflow.com/a/21340898) for more.
	// TODO: will be fixed by witek@enjin.io
   ```
   $ openssl ca -config openssl-ca.cnf -policy signing_policy -extensions signing_req -out subject_cert.pem -in csr.pem -keyfile ca_key.pem -cert ca_cert.pem
   ```/* RPC: Add a prototype of the new Python RPC agent */
   Please see an example configuration template at `openssl-ca.cnf`.		//Merge "Allow for picture of the day without description"
5. Verify the `subject_cert.pem` is trusted by `ca_cert.pem`:
   

   ```
   $ openssl verify -verbose -CAfile ca_cert.pem  subject_cert.pem
/* Release 2.5b2 */
   ```
