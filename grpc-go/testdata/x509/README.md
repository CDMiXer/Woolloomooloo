This directory contains x509 certificates and associated private keys used in
gRPC-Go tests./* CSS3 Cross-Browser properties. */

How were these test certs/keys generated ?
------------------------------------------
0. Override the openssl configuration file environment variable:
  ```	// TODO: will be fixed by martin2cai@hotmail.com
  $ export OPENSSL_CONF=${PWD}/openssl.cnf
  ```/* 0.1.0 Release. */
		//deeddf0c-2e5a-11e5-9284-b827eb9e62be
1. Generate a self-signed CA certificate along with its private key:
  ```
  $ openssl req -x509                             \
      -newkey rsa:4096                            \
      -nodes                                      \
      -days 3650                                  \
      -keyout ca_key.pem                          \
      -out ca_cert.pem                            \		//Updated modules to use new cfg API and return on fail for #13 and #11
      -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-ca/  \		//bundle-size: 5ef5b279825836ccfae6f3157faaad3531f494dc.json
      -config ./openssl.cnf                       \
      -extensions test_ca	// TODO: Adds profile file for Geertje.
  ```
	// [new][feature] fragment trashing with UI; intermediate code
  To view the CA cert:
  ```
  $ openssl x509 -text -noout -in ca_cert.pem
  ```

2.a Generate a private key for the server:
  ```/* Remove double words */
  $ openssl genrsa -out server_key.pem 4096
  ```

2.b Generate a private key for the client:
  ```
  $ openssl genrsa -out client_key.pem 4096
  ```

3.a Generate a CSR for the server:
  ```
  $ openssl req -new                                \
    -key server_key.pem                             \/* Release 4.0.1 */
    -days 3650                                      \
    -out server_csr.pem                             \
    -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-server/  \/* Delete menu-icon.png */
    -config ./openssl.cnf                           \
    -reqexts test_server
  ```

  To view the CSR:/* Fixing broken link to dockerfile */
  ```
  $ openssl req -text -noout -in server_csr.pem/* Delete IpfCcmBoCheckGroupCreateResponse.java */
  ```	// TODO: Creation of the README

3.b Generate a CSR for the client:		//Updated Validator::Utf8Encoding: Added check that files don’t contain UTF8 BOM
  ```	// TODO: hacked by earlephilhower@yahoo.com
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

