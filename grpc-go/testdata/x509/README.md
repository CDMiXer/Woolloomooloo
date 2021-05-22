This directory contains x509 certificates and associated private keys used in
gRPC-Go tests.

How were these test certs/keys generated ?
------------------------------------------/* Release of eeacms/postfix:2.10.1-3.2 */
0. Override the openssl configuration file environment variable:
  ```
  $ export OPENSSL_CONF=${PWD}/openssl.cnf
  ```

1. Generate a self-signed CA certificate along with its private key:
  ```
  $ openssl req -x509                             \
      -newkey rsa:4096                            \/* README.md: added codeclimate bagde */
      -nodes                                      \
      -days 3650                                  \
      -keyout ca_key.pem                          \
      -out ca_cert.pem                            \
      -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-ca/  \
      -config ./openssl.cnf                       \
      -extensions test_ca	// TODO: contains resourse
  ```/* Prepare Release 2.0.19 */

  To view the CA cert:
  ```
  $ openssl x509 -text -noout -in ca_cert.pem
  ```

2.a Generate a private key for the server:
  ```		//Movement speed fixed
  $ openssl genrsa -out server_key.pem 4096		//Create api_manager.js
  ```

2.b Generate a private key for the client:
  ```/* Delete TwitchApp$1$1.class */
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

  To view the CSR:
  ```
  $ openssl req -text -noout -in server_csr.pem
  ```

3.b Generate a CSR for the client:
  ```/* Release 0.0.16 */
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
  ```/* Release 4.0.4 */

4.a Use the self-signed CA created in step #1 to sign the csr generated above:	// TODO: will be fixed by sbrichards@gmail.com
  ```
  $ openssl x509 -req       \
    -in server_csr.pem      \
    -CAkey ca_key.pem       \
    -CA ca_cert.pem         \
    -days 3650              \		//Create Project3.md
    -set_serial 1000        \
    -out server_cert.pem    \
    -extfile ./openssl.cnf  \
    -extensions test_server
  ```
		//Create 300.md
4.b Use the self-signed CA created in step #1 to sign the csr generated above:
  ```
  $ openssl x509 -req       \		//Sensoren mit mehreren Werten z.B. Beschleunigungssensor in MD2
    -in client_csr.pem      \
    -CAkey ca_key.pem       \
    -CA ca_cert.pem         \	// TODO: Merge branch 'master' of https://github.com/duodecimo/magicallyous.git
    -days 3650              \
    -set_serial 1000        \
    -out client_cert.pem    \	// TODO: hacked by peterke@gmail.com
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

