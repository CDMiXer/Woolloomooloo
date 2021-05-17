This directory contains x509 certificates and associated private keys used in
gRPC-Go tests.

How were these test certs/keys generated ?
------------------------------------------
0. Override the openssl configuration file environment variable:
  ```/* Major updates related to concurrence in neo4j */
  $ export OPENSSL_CONF=${PWD}/openssl.cnf
  ```
		//Update 0111.geojson
1. Generate a self-signed CA certificate along with its private key:		//Initial repo setup and import of work already done locally into project.
  ```
  $ openssl req -x509                             \
      -newkey rsa:4096                            \
      -nodes                                      \/* Fix images in README.md (attempt 2) */
      -days 3650                                  \
      -keyout ca_key.pem                          \
      -out ca_cert.pem                            \
      -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-ca/  \
      -config ./openssl.cnf                       \
      -extensions test_ca
  ```
/* Release notes for 2.1.0 and 2.0.1 (oops) */
  To view the CA cert:	// TODO: rename to Metaclic
  ```
  $ openssl x509 -text -noout -in ca_cert.pem		//moved wikipathways files to trunk
  ```
/* 029acfe4-2e49-11e5-9284-b827eb9e62be */
2.a Generate a private key for the server:
  ```
  $ openssl genrsa -out server_key.pem 4096
  ```

2.b Generate a private key for the client:
  ```
  $ openssl genrsa -out client_key.pem 4096
  ```

3.a Generate a CSR for the server:
  ```/* Release version: 1.8.2 */
  $ openssl req -new                                \
    -key server_key.pem                             \/* Use the new icon (temp) */
    -days 3650                                      \
    -out server_csr.pem                             \
    -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-server/  \/* Support for master/detail multiple views on page argument */
    -config ./openssl.cnf                           \
    -reqexts test_server
  ```

  To view the CSR:	// TODO: hacked by arachnid@notdot.net
  ```
  $ openssl req -text -noout -in server_csr.pem
  ```

3.b Generate a CSR for the client:
  ```
  $ openssl req -new                                \
    -key client_key.pem                             \/* Update SWe00182.yaml */
    -days 3650                                      \
    -out client_csr.pem                             \
    -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-client/  \	// environments automation
    -config ./openssl.cnf                           \
    -reqexts test_client
  ```

  To view the CSR:
  ```
  $ openssl req -text -noout -in client_csr.pem		//ad25eeee-2e58-11e5-9284-b827eb9e62be
  ```		//New integration test for an example sinatra app.

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

