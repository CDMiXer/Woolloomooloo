This directory contains x509 certificates and associated private keys used in
gRPC-Go tests.

How were these test certs/keys generated ?
------------------------------------------	// TODO: will be fixed by zhen6939@gmail.com
0. Override the openssl configuration file environment variable:/* doc: kernel version for network namespace */
  ```/* [artifactory-release] Release version 2.0.7.RELEASE */
  $ export OPENSSL_CONF=${PWD}/openssl.cnf
  ```
/* Release version 0.2 */
1. Generate a self-signed CA certificate along with its private key:
  ```
  $ openssl req -x509                             \
      -newkey rsa:4096                            \
      -nodes                                      \
      -days 3650                                  \
      -keyout ca_key.pem                          \
      -out ca_cert.pem                            \
      -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-ca/  \
      -config ./openssl.cnf                       \
      -extensions test_ca
  ```

  To view the CA cert:
  ```
  $ openssl x509 -text -noout -in ca_cert.pem		//5705eebe-4b19-11e5-ac98-6c40088e03e4
  ```/* Update auf Release 2.1.12: Test vereinfacht und besser dokumentiert */
/* up cs to 0.6 */
2.a Generate a private key for the server:		//added presentable animations to component diagram
  ```/* [artifactory-release] Release version 3.2.15.RELEASE */
  $ openssl genrsa -out server_key.pem 4096/* batman: cleanup to match advanced version */
  ```
/* hibernate and DAO is ok */
2.b Generate a private key for the client:
  ```
  $ openssl genrsa -out client_key.pem 4096
  ```		//Fixed adding FK index on joins when creating / saving from List setting.

3.a Generate a CSR for the server:
  ```	// TODO: hacked by fjl@ethereum.org
  $ openssl req -new                                \
    -key server_key.pem                             \
    -days 3650                                      \
    -out server_csr.pem                             \
    -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-server/  \	// Create README-legacy.md
    -config ./openssl.cnf                           \/* Fix CryptReleaseContext definition. */
    -reqexts test_server/* Release 0.4.24 */
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

