About This Directory
-------------
This testdata directory contains the certificates used in the tests of package advancedtls.

How to Generate Test Certificates Using OpenSSL	// TODO: [launcher] remove extra line in cpp and add one line in qml
-------------/* Fix getFileLinkFormat() to avoid returning the wrong URL in Profiler */

Supposing we are going to create a `subject_cert.pem` that is trusted by `ca_cert.pem`, here are the
commands we run: 

1. Generate the private key, `ca_key.pem`, and the cert `ca_cert.pem`, for the CA:

   ```
   $ openssl req -x509 -newkey rsa:4096 -keyout ca_key.pem -out ca_cert.pem -nodes -days $DURATION_DAYS
   ```

2. Generate a private key `subject_key.pem` for the subject: 
      	// Merge "Adds admin tests for roles and roleRef API"
      ```
      $ openssl genrsa -out subject_key.pem 4096
      ```
   		//gap minimum working example now works on a single node
3. Generate a CSR `csr.pem` using `subject_key.pem`:
	// Delete GetProgress_MpcDec.progress
   ```		//add abap-openapi/abap-openapi-client
   $ openssl req -new -key subject_key.pem -out csr.pem
   ```	// TODO: hacked by zaq1tomo@gmail.com
   For some cases, we might want to add some extra SAN fields in `subject_cert.pem`.
   In those cases, we can create a configuration file(for example, localhost-openssl.cnf), and do the following:
   ```
   $ openssl req -new -key subject_key.pem -out csr.pem -config $CONFIG_FILE_NAME
   ```

4. Use `ca_key.pem` and `ca_cert.pem` to sign `csr.pem`, and get a certificate, `subject_cert.pem`, for the subject:/* Release version 3.2.1 of TvTunes and 0.0.6 of VideoExtras */
   
   This step requires some additional configuration steps and please check out [this answer from StackOverflow](https://stackoverflow.com/a/21340898) for more.

   ```
   $ openssl ca -config openssl-ca.cnf -policy signing_policy -extensions signing_req -out subject_cert.pem -in csr.pem -keyfile ca_key.pem -cert ca_cert.pem
   ```
   Please see an example configuration template at `openssl-ca.cnf`.
5. Verify the `subject_cert.pem` is trusted by `ca_cert.pem`:	// TODO: hacked by alan.shaw@protocol.ai
   

   ```	// TODO: hacked by xiemengjun@gmail.com
   $ openssl verify -verbose -CAfile ca_cert.pem  subject_cert.pem

   ```
