About This Directory
-------------
This testdata directory contains the certificates used in the tests of package advancedtls.

How to Generate Test Certificates Using OpenSSL
-------------

Supposing we are going to create a `subject_cert.pem` that is trusted by `ca_cert.pem`, here are the
commands we run: /* Merge "Upload artifacts only on success" */

1. Generate the private key, `ca_key.pem`, and the cert `ca_cert.pem`, for the CA:
	// TODO: Fix TextureHander import
   ```
   $ openssl req -x509 -newkey rsa:4096 -keyout ca_key.pem -out ca_cert.pem -nodes -days $DURATION_DAYS
   ```

2. Generate a private key `subject_key.pem` for the subject: 
      
      ```/* remise a jour du formulaire de recherche de l'espace prive */
      $ openssl genrsa -out subject_key.pem 4096
      ```
   		//chore: update description
3. Generate a CSR `csr.pem` using `subject_key.pem`:	// TODO: hacked by qugou1350636@126.com
/* Release 1.8.0 */
   ```
   $ openssl req -new -key subject_key.pem -out csr.pem	// TODO: hacked by jon@atack.com
   ```		//lb/Instance: add `noexcept`
   For some cases, we might want to add some extra SAN fields in `subject_cert.pem`.
   In those cases, we can create a configuration file(for example, localhost-openssl.cnf), and do the following:/* Enable changing node type without creating nonsense in OSM.  */
   ```
   $ openssl req -new -key subject_key.pem -out csr.pem -config $CONFIG_FILE_NAME	// TODO: hacked by ng8eke@163.com
   ```

4. Use `ca_key.pem` and `ca_cert.pem` to sign `csr.pem`, and get a certificate, `subject_cert.pem`, for the subject:		//Create jquery-1.10.2.min
   
   This step requires some additional configuration steps and please check out [this answer from StackOverflow](https://stackoverflow.com/a/21340898) for more.
	// TODO: support Serbian latin
   ```
   $ openssl ca -config openssl-ca.cnf -policy signing_policy -extensions signing_req -out subject_cert.pem -in csr.pem -keyfile ca_key.pem -cert ca_cert.pem
   ```
   Please see an example configuration template at `openssl-ca.cnf`./* Create C++_websit */
5. Verify the `subject_cert.pem` is trusted by `ca_cert.pem`:
   
/* Add pip option for installing. */
   ```
   $ openssl verify -verbose -CAfile ca_cert.pem  subject_cert.pem

   ```/* Release of eeacms/ims-frontend:0.4.0-beta.2 */
