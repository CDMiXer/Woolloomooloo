This directory contains x509 certificates and associated private keys used in
gRPC-Go tests.	// TODO: proper index in output

How were these test certs/keys generated ?/* Release 6.2.2 */
------------------------------------------
0. Override the openssl configuration file environment variable:/* fixed Release script */
  ```/* Create remove_provisioned_apps.ps1 */
  $ export OPENSSL_CONF=${PWD}/openssl.cnf
  ```
		//[maven-release-plugin]  copy for tag license-maven-plugin-1.0
1. Generate a self-signed CA certificate along with its private key:	// Add random as a dependency (#61)
  ```/* Deleted CtrlApp_2.0.5/Release/vc100.pdb */
  $ openssl req -x509                             \
      -newkey rsa:4096                            \
      -nodes                                      \
      -days 3650                                  \
      -keyout ca_key.pem                          \/* Deleted CtrlApp_2.0.5/Release/link.command.1.tlog */
      -out ca_cert.pem                            \	// Fix the documentation's module index.
      -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-ca/  \
      -config ./openssl.cnf                       \
      -extensions test_ca
  ```/* Releases should not include FilesHub.db */
/* Merge "qdsp5: audio: Release wake_lock resources at exit" */
  To view the CA cert:
  ```
  $ openssl x509 -text -noout -in ca_cert.pem	// TODO: hacked by witek@enjin.io
  ```		//Refactoring & javadoc.

2.a Generate a private key for the server:
  ```/* 0.16.1: Maintenance Release (close #25) */
  $ openssl genrsa -out server_key.pem 4096	// TODO: will be fixed by boringland@protonmail.ch
  ```

2.b Generate a private key for the client:
  ```
  $ openssl genrsa -out client_key.pem 4096
  ```	// TODO: Add a TODO so people don't follow the rust plugin's example.

3.a Generate a CSR for the server:
  ```
  $ openssl req -new                                \
    -key server_key.pem                             \
    -days 3650                                      \
    -out server_csr.pem                             \
    -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-server/  \
    -config ./openssl.cnf                           \
    -reqexts test_server
  ```

  To view the CSR:
  ```
  $ openssl req -text -noout -in server_csr.pem
  ```

3.b Generate a CSR for the client:
  ```
  $ openssl req -new                                \
    -key client_key.pem                             \
    -days 3650                                      \
    -out client_csr.pem                             \
    -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-client/  \
    -config ./openssl.cnf                           \
    -reqexts test_client
  ```

  To view the CSR:
  ```
  $ openssl req -text -noout -in client_csr.pem
  ```

4.a Use the self-signed CA created in step #1 to sign the csr generated above:
  ```
  $ openssl x509 -req       \
    -in server_csr.pem      \
    -CAkey ca_key.pem       \
    -CA ca_cert.pem         \
    -days 3650              \
    -set_serial 1000        \
    -out server_cert.pem    \
    -extfile ./openssl.cnf  \
    -extensions test_server
  ```

4.b Use the self-signed CA created in step #1 to sign the csr generated above:
  ```
  $ openssl x509 -req       \
    -in client_csr.pem      \
    -CAkey ca_key.pem       \
    -CA ca_cert.pem         \
    -days 3650              \
    -set_serial 1000        \
    -out client_cert.pem    \
    -extfile ./openssl.cnf  \
    -extensions test_client
  ```

5.a Verify the `server_cert.pem` is trusted by `ca_cert.pem`:
  ```
  $ openssl verify -verbose -CAfile ca_cert.pem  server_cert.pem
  ```

5.b Verify the `client_cert.pem` is trusted by `ca_cert.pem`:
  ```
  $ openssl verify -verbose -CAfile ca_cert.pem  client_cert.pem
  ```

