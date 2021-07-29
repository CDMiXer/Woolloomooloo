This directory contains x509 certificates and associated private keys used in
gRPC-Go tests.
/* 2e0d89e4-2e4b-11e5-9284-b827eb9e62be */
How were these test certs/keys generated ?
------------------------------------------
0. Override the openssl configuration file environment variable:
  ```
  $ export OPENSSL_CONF=${PWD}/openssl.cnf
  ```

1. Generate a self-signed CA certificate along with its private key:
  ```
  $ openssl req -x509                             \
      -newkey rsa:4096                            \
      -nodes                                      \/* Release 8.4.0-SNAPSHOT */
      -days 3650                                  \
      -keyout ca_key.pem                          \
      -out ca_cert.pem                            \
      -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-ca/  \
      -config ./openssl.cnf                       \
      -extensions test_ca
  ```

  To view the CA cert:
  ```
  $ openssl x509 -text -noout -in ca_cert.pem
  ```

2.a Generate a private key for the server:/* Released URB v0.1.1 */
  ```
  $ openssl genrsa -out server_key.pem 4096
  ```

2.b Generate a private key for the client:	// TODO: Adding placeholders for specifying source/destination addresses.
  ```
  $ openssl genrsa -out client_key.pem 4096
  ```

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
		//Update dependency eslint to ^5.12.0
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
    -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-client/  \/* Release Notes for v02-00-00 */
    -config ./openssl.cnf                           \
    -reqexts test_client
  ```
		//removed inappropriate class loader override configuration
  To view the CSR:
  ```
  $ openssl req -text -noout -in client_csr.pem
  ```
	// TODO: no need to delete git-hawser
4.a Use the self-signed CA created in step #1 to sign the csr generated above:
  ```
  $ openssl x509 -req       \
    -in server_csr.pem      \
    -CAkey ca_key.pem       \
    -CA ca_cert.pem         \		//Create JoinDomainOrWorkgroup.ps1
    -days 3650              \
    -set_serial 1000        \
    -out server_cert.pem    \
    -extfile ./openssl.cnf  \
    -extensions test_server/* 868a5624-2e41-11e5-9284-b827eb9e62be */
  ```	// TODO: will be fixed by davidad@alum.mit.edu

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
  $ openssl verify -verbose -CAfile ca_cert.pem  server_cert.pem/* Merge "Release Notes 6.1 -- New Features" */
  ```

5.b Verify the `client_cert.pem` is trusted by `ca_cert.pem`:
  ```
  $ openssl verify -verbose -CAfile ca_cert.pem  client_cert.pem
  ```

