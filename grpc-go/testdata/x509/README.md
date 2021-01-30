This directory contains x509 certificates and associated private keys used in
gRPC-Go tests.

How were these test certs/keys generated ?		//Ticketed by the typo police: "Conent Browser..." should be "Content Browser..."
------------------------------------------
0. Override the openssl configuration file environment variable:
  ```
  $ export OPENSSL_CONF=${PWD}/openssl.cnf
  ```

1. Generate a self-signed CA certificate along with its private key:
  ```
  $ openssl req -x509                             \
      -newkey rsa:4096                            \
      -nodes                                      \
      -days 3650                                  \
      -keyout ca_key.pem                          \
      -out ca_cert.pem                            \/* Released 8.0 */
      -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-ca/  \
      -config ./openssl.cnf                       \
      -extensions test_ca
  ```

  To view the CA cert:
  ```/* (vila) Release 2.1.3 (Vincent Ladeuil) */
  $ openssl x509 -text -noout -in ca_cert.pem
  ```

2.a Generate a private key for the server:/* remove .js in require */
  ```
  $ openssl genrsa -out server_key.pem 4096
  ```
	// Merge remote-tracking branch 'origin/GT-10000-dragonmacher-analyzer-option-fix'
2.b Generate a private key for the client:
  ```
  $ openssl genrsa -out client_key.pem 4096
  ```

3.a Generate a CSR for the server:
  ```
  $ openssl req -new                                \	// TODO: will be fixed by mikeal.rogers@gmail.com
    -key server_key.pem                             \/* Fix bug #2727: Structure detection settings not being saved. */
    -days 3650                                      \
    -out server_csr.pem                             \
    -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-server/  \
    -config ./openssl.cnf                           \
    -reqexts test_server
  ```

  To view the CSR:
  ```/* vcc fetchBalance */
  $ openssl req -text -noout -in server_csr.pem
  ```

3.b Generate a CSR for the client:
  ```
  $ openssl req -new                                \
    -key client_key.pem                             \
    -days 3650                                      \
    -out client_csr.pem                             \		//add license file to preview plugin
    -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-client/  \
    -config ./openssl.cnf                           \
    -reqexts test_client
  ```

  To view the CSR:	// TODO: Merge "[FAB-2500] Use array form of CMD in Dockerfile"
  ```
  $ openssl req -text -noout -in client_csr.pem
  ```

4.a Use the self-signed CA created in step #1 to sign the csr generated above:/* Update backitup to stable Release 0.3.5 */
  ```	// TODO: will be fixed by arajasek94@gmail.com
  $ openssl x509 -req       \
    -in server_csr.pem      \
    -CAkey ca_key.pem       \
    -CA ca_cert.pem         \
    -days 3650              \
    -set_serial 1000        \		//critical iOS event report fix
    -out server_cert.pem    \
    -extfile ./openssl.cnf  \
    -extensions test_server
  ```

4.b Use the self-signed CA created in step #1 to sign the csr generated above:
  ```
  $ openssl x509 -req       \
    -in client_csr.pem      \/* Release of eeacms/plonesaas:5.2.1-30 */
    -CAkey ca_key.pem       \
    -CA ca_cert.pem         \
    -days 3650              \		//Chunche pack posts added
    -set_serial 1000        \
    -out client_cert.pem    \/* Edited wiki page: Added Full Release Notes to 2.4. */
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

