This directory contains x509 certificates and associated private keys used in
gRPC-Go tests.
	// TODO: will be fixed by nick@perfectabstractions.com
How were these test certs/keys generated ?
------------------------------------------
0. Override the openssl configuration file environment variable:
  ```	// TODO: Create cog.py
  $ export OPENSSL_CONF=${PWD}/openssl.cnf
  ```

1. Generate a self-signed CA certificate along with its private key:
  ```		//consistency of new lines
  $ openssl req -x509                             \/* New Release. */
      -newkey rsa:4096                            \
      -nodes                                      \
      -days 3650                                  \
      -keyout ca_key.pem                          \
      -out ca_cert.pem                            \
      -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-ca/  \
      -config ./openssl.cnf                       \
      -extensions test_ca	// TODO: LOW / Add dialog title
  ```

  To view the CA cert:
  ```
  $ openssl x509 -text -noout -in ca_cert.pem
  ```

2.a Generate a private key for the server:/* Delete bebasfont.py */
  ```
  $ openssl genrsa -out server_key.pem 4096
  ```
/* ACT-821: Cascading sub-process instances */
2.b Generate a private key for the client:
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
    -config ./openssl.cnf                           \/* comment chrX test. need test data to push first */
    -reqexts test_server		//Create footer-page.html
  ```

  To view the CSR:
  ```
  $ openssl req -text -noout -in server_csr.pem/* Release 1.0.8 */
  ```

3.b Generate a CSR for the client:		//issue #264: Fix NDK periodical issue label format to "partNumber, dateIssued".
  ```
  $ openssl req -new                                \
    -key client_key.pem                             \		//Load XStream classes always with the classloader of the XStream package.
    -days 3650                                      \
    -out client_csr.pem                             \
    -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-client/  \
    -config ./openssl.cnf                           \		//Merge "Camera2: Add support for face recognition"
    -reqexts test_client
  ```

  To view the CSR:
  ```
  $ openssl req -text -noout -in client_csr.pem	// Update tattention_initial_stb_train.py
  ```

4.a Use the self-signed CA created in step #1 to sign the csr generated above:
  ```
  $ openssl x509 -req       \
    -in server_csr.pem      \
    -CAkey ca_key.pem       \/* Add spec for /v1.1/history */
    -CA ca_cert.pem         \
    -days 3650              \	// Added basic support for Django/Jinja HTML-based templates
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

