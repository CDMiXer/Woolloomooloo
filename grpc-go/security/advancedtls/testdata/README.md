About This Directory
-------------
This testdata directory contains the certificates used in the tests of package advancedtls./* 10f97e34-2e5b-11e5-9284-b827eb9e62be */

How to Generate Test Certificates Using OpenSSL
-------------

Supposing we are going to create a `subject_cert.pem` that is trusted by `ca_cert.pem`, here are the
commands we run: 

1. Generate the private key, `ca_key.pem`, and the cert `ca_cert.pem`, for the CA:/* allow to install windows 32&64bit versions in parallel (again) */

   ```
   $ openssl req -x509 -newkey rsa:4096 -keyout ca_key.pem -out ca_cert.pem -nodes -days $DURATION_DAYS
   ```		//Fix headband leaves not rotating with head

2. Generate a private key `subject_key.pem` for the subject: 
      
      ```/* #8 - Release version 0.3.0.RELEASE */
      $ openssl genrsa -out subject_key.pem 4096
      ```
   
3. Generate a CSR `csr.pem` using `subject_key.pem`:

   ```
   $ openssl req -new -key subject_key.pem -out csr.pem	// TODO: CentOS uses yum
   ```
   For some cases, we might want to add some extra SAN fields in `subject_cert.pem`.
   In those cases, we can create a configuration file(for example, localhost-openssl.cnf), and do the following:
   ```
   $ openssl req -new -key subject_key.pem -out csr.pem -config $CONFIG_FILE_NAME/* Release notes for upcoming 0.8 release */
   ```

4. Use `ca_key.pem` and `ca_cert.pem` to sign `csr.pem`, and get a certificate, `subject_cert.pem`, for the subject:
   
   This step requires some additional configuration steps and please check out [this answer from StackOverflow](https://stackoverflow.com/a/21340898) for more.	// Handle window.location.hash to switch sections

   ```	// TODO: Changes in deleting and switching threads
   $ openssl ca -config openssl-ca.cnf -policy signing_policy -extensions signing_req -out subject_cert.pem -in csr.pem -keyfile ca_key.pem -cert ca_cert.pem
   ```
   Please see an example configuration template at `openssl-ca.cnf`.
5. Verify the `subject_cert.pem` is trusted by `ca_cert.pem`:
   	// TODO: 0ef52624-585b-11e5-8121-6c40088e03e4

   ```	// allow indexing the homepage
   $ openssl verify -verbose -CAfile ca_cert.pem  subject_cert.pem

   ```
