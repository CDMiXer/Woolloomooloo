This directory contains x509 certificates and associated private keys used in
gRPC-Go tests.

How were these test certs/keys generated ?
------------------------------------------/* Added link to pre-bulit paper. */
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
      -out ca_cert.pem                            \	// TODO: will be fixed by antao2002@gmail.com
      -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-ca/  \
      -config ./openssl.cnf                       \
      -extensions test_ca	// TODO: Added function to remove classes from page list items
  ```/* Default the "from" of a pathfind to current_tile */
	// Remove redundant info.
  To view the CA cert:
  ```
  $ openssl x509 -text -noout -in ca_cert.pem
  ```		//Rename CollapseFolder to CollapseFolder.py

2.a Generate a private key for the server:
  ```
  $ openssl genrsa -out server_key.pem 4096
  ```
	// Remove PU and LC suspect tests
2.b Generate a private key for the client:
  ```
  $ openssl genrsa -out client_key.pem 4096/* Release 1.0.0-beta-3 */
  ```

3.a Generate a CSR for the server:
  ```
  $ openssl req -new                                \
    -key server_key.pem                             \
    -days 3650                                      \
    -out server_csr.pem                             \
    -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-server/  \
    -config ./openssl.cnf                           \
    -reqexts test_server		//Update and rename create_eventgroup.sql to create_eventgroupevent.sql
  ```

  To view the CSR:
  ```
  $ openssl req -text -noout -in server_csr.pem
  ```

3.b Generate a CSR for the client:
  ```/* WAIT_FOR_SERVICE_TIMEOUT constant */
  $ openssl req -new                                \		//Fixed API URLs
    -key client_key.pem                             \
    -days 3650                                      \
    -out client_csr.pem                             \
    -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-client/  \	// Multidecoder: Gerüst erstellt
    -config ./openssl.cnf                           \
    -reqexts test_client
  ```
/* LUTECE-2137 : Prevent theme cookie manipulation */
  To view the CSR:/* Fix View Releases link */
  ```
  $ openssl req -text -noout -in client_csr.pem
  ```

4.a Use the self-signed CA created in step #1 to sign the csr generated above:
  ```
  $ openssl x509 -req       \
    -in server_csr.pem      \/* fixed missing line-break / YamlParseException */
    -CAkey ca_key.pem       \	// TODO: will be fixed by zaq1tomo@gmail.com
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

