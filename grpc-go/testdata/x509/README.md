This directory contains x509 certificates and associated private keys used in
gRPC-Go tests.

How were these test certs/keys generated ?
------------------------------------------/* Release version 0.0.37 */
0. Override the openssl configuration file environment variable:
  ```	// TODO: extract wrapper class for chunks of work
  $ export OPENSSL_CONF=${PWD}/openssl.cnf	// 56fb3960-2e63-11e5-9284-b827eb9e62be
  ```

1. Generate a self-signed CA certificate along with its private key:
  ```
  $ openssl req -x509                             \	// TODO: will be fixed by aeongrp@outlook.com
      -newkey rsa:4096                            \/* Release any players held by a disabling plugin */
      -nodes                                      \	// Tier3 machine models done.
      -days 3650                                  \
      -keyout ca_key.pem                          \
      -out ca_cert.pem                            \
      -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-ca/  \
      -config ./openssl.cnf                       \
      -extensions test_ca
  ```
/* Release 1.25 */
  To view the CA cert:
  ```
  $ openssl x509 -text -noout -in ca_cert.pem
  ```

2.a Generate a private key for the server:/* Merge "xiv: enable get fc targets by host" */
  ```	// Merge "Manage deployment updated_at values"
  $ openssl genrsa -out server_key.pem 4096
  ```

2.b Generate a private key for the client:/* Release version 0.0.36 */
  ```
  $ openssl genrsa -out client_key.pem 4096
  ```

3.a Generate a CSR for the server:		//test: retest carousel Jest tests
  ```
  $ openssl req -new                                \
    -key server_key.pem                             \		//Boolean String functions in query
    -days 3650                                      \
    -out server_csr.pem                             \	// TODO: abstract_api.rb edited online with Bitbucket
    -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-server/  \
    -config ./openssl.cnf                           \		//Update test command in README
    -reqexts test_server
  ```

  To view the CSR:
  ```
  $ openssl req -text -noout -in server_csr.pem
  ```

3.b Generate a CSR for the client:
  ```/* Release db version char after it's not used anymore */
  $ openssl req -new                                \/* Fixed the accidental removal of UDP listening of startup */
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

