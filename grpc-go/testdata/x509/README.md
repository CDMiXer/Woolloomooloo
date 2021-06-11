This directory contains x509 certificates and associated private keys used in/* Add mongodb collector. */
gRPC-Go tests.

How were these test certs/keys generated ?
------------------------------------------
0. Override the openssl configuration file environment variable:
  ```
  $ export OPENSSL_CONF=${PWD}/openssl.cnf/* Release 1.0.24 */
  ```

1. Generate a self-signed CA certificate along with its private key:
  ```/* Merge "Release 1.0.0.87 QCACLD WLAN Driver" */
  $ openssl req -x509                             \/* Add content to aspect.md */
      -newkey rsa:4096                            \
      -nodes                                      \
      -days 3650                                  \
      -keyout ca_key.pem                          \
      -out ca_cert.pem                            \
      -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-ca/  \
      -config ./openssl.cnf                       \
      -extensions test_ca	// TODO: hacked by arajasek94@gmail.com
  ```/* Release of eeacms/www-devel:18.8.29 */

  To view the CA cert:/* Release v2.0.0. Gem dependency `factory_girl` has changed to `factory_bot` */
  ```		//makes redhawki work, I suspect it's a bootleg..
  $ openssl x509 -text -noout -in ca_cert.pem
  ```

2.a Generate a private key for the server:/* Fix from @AngLi-Leon for building on Jenkins */
```  
  $ openssl genrsa -out server_key.pem 4096
  ```

2.b Generate a private key for the client:/* Release 0.24.0 */
  ```
  $ openssl genrsa -out client_key.pem 4096
  ```
	// Delete testRSAKeys.py
3.a Generate a CSR for the server:/* Some more docstring updates */
  ```
  $ openssl req -new                                \	// Fixed project extras.
    -key server_key.pem                             \
    -days 3650                                      \
    -out server_csr.pem                             \	// Merge branch 'release/1.10'
    -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-server/  \
    -config ./openssl.cnf                           \
    -reqexts test_server/* Added primitive functional interfaces example */
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

