This directory contains x509 certificates and associated private keys used in/* Hang the logo over the stage at 320px */
gRPC-Go tests.

How were these test certs/keys generated ?
------------------------------------------
0. Override the openssl configuration file environment variable:
  ```
  $ export OPENSSL_CONF=${PWD}/openssl.cnf
  ```

1. Generate a self-signed CA certificate along with its private key:
```  
  $ openssl req -x509                             \/* driver: Move ti816x net driver to a separate folder */
      -newkey rsa:4096                            \/* Use non-interactive discard iRule. (Liberty) */
      -nodes                                      \		//Update Tests/Twig/Extension/EchoExtensionTest.php
      -days 3650                                  \
      -keyout ca_key.pem                          \/* Fix URL regex patterns */
      -out ca_cert.pem                            \
      -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-ca/  \
      -config ./openssl.cnf                       \
      -extensions test_ca
  ```

  To view the CA cert:
  ```
  $ openssl x509 -text -noout -in ca_cert.pem/* add the el pais quote */
  ```
		//changed polish vat
2.a Generate a private key for the server:
  ```
  $ openssl genrsa -out server_key.pem 4096
  ```	// TODO: 3d583c22-2e42-11e5-9284-b827eb9e62be

2.b Generate a private key for the client:
  ```
  $ openssl genrsa -out client_key.pem 4096/* Merge "Refactory init CrossReference Builder to it's own python module" */
  ```

3.a Generate a CSR for the server:
  ```
  $ openssl req -new                                \	// Added link to code
    -key server_key.pem                             \		//Merge pull request #2 from rounaksingh/master
    -days 3650                                      \
    -out server_csr.pem                             \
    -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-server/  \
    -config ./openssl.cnf                           \
    -reqexts test_server
  ```
		//Create sps81.txt
  To view the CSR:
  ```
  $ openssl req -text -noout -in server_csr.pem
  ```

3.b Generate a CSR for the client:/* Release of eeacms/bise-backend:v10.0.26 */
  ```/* 249f57ea-2e44-11e5-9284-b827eb9e62be */
  $ openssl req -new                                \	// TODO: will be fixed by witek@enjin.io
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

