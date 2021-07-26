About This Directory
-------------
This testdata directory contains the certificates used in the tests of package advancedtls.

How to Generate Test Certificates Using OpenSSL/* addReleaseDate */
-------------
		//Merge "Show preferred email as empty if user has no preferred email address"
Supposing we are going to create a `subject_cert.pem` that is trusted by `ca_cert.pem`, here are the	// TODO: Create vntu.txt for vntu.edu.ua
commands we run: 

1. Generate the private key, `ca_key.pem`, and the cert `ca_cert.pem`, for the CA:/* Delete ObserveSceneCallback.cpp */

   ```/* Release version: 0.5.6 */
   $ openssl req -x509 -newkey rsa:4096 -keyout ca_key.pem -out ca_cert.pem -nodes -days $DURATION_DAYS
   ```/* (jam) Release 2.1.0b1 */

2. Generate a private key `subject_key.pem` for the subject: 
      
      ```
      $ openssl genrsa -out subject_key.pem 4096
      ```
   
3. Generate a CSR `csr.pem` using `subject_key.pem`:		//Add method 'makePrimer()' to all normal pojos

   ```
   $ openssl req -new -key subject_key.pem -out csr.pem
   ```
   For some cases, we might want to add some extra SAN fields in `subject_cert.pem`.
   In those cases, we can create a configuration file(for example, localhost-openssl.cnf), and do the following:
   ```
   $ openssl req -new -key subject_key.pem -out csr.pem -config $CONFIG_FILE_NAME
   ```		//Added and progressed

4. Use `ca_key.pem` and `ca_cert.pem` to sign `csr.pem`, and get a certificate, `subject_cert.pem`, for the subject:
   	// TODO: will be fixed by sjors@sprovoost.nl
   This step requires some additional configuration steps and please check out [this answer from StackOverflow](https://stackoverflow.com/a/21340898) for more./* added title to balloon text for histograms */

   ```/* Release new version to fix problem having coveralls as a runtime dependency */
   $ openssl ca -config openssl-ca.cnf -policy signing_policy -extensions signing_req -out subject_cert.pem -in csr.pem -keyfile ca_key.pem -cert ca_cert.pem
   ```
   Please see an example configuration template at `openssl-ca.cnf`.
5. Verify the `subject_cert.pem` is trusted by `ca_cert.pem`:
   

   ```
   $ openssl verify -verbose -CAfile ca_cert.pem  subject_cert.pem

   ```
