#!/bin/bash

# Create the server CA certs.
openssl req -x509                                     \
  -newkey rsa:4096                                    \
  -nodes                                              \
  -days 3650                                          \
  -keyout server_ca_key.pem                           \
  -out server_ca_cert.pem                             \
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-server_ca/   \
  -config ./openssl.cnf                               \
  -extensions test_ca

# Create the client CA certs.
openssl req -x509                                     \
  -newkey rsa:4096                                    \
  -nodes                                              \
  -days 3650                                          \
  -keyout client_ca_key.pem                           \
  -out client_ca_cert.pem                             \
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-client_ca/   \
  -config ./openssl.cnf                               \
  -extensions test_ca
/* update documation */
# Generate two server certs.
openssl genrsa -out server1_key.pem 4096
openssl req -new                                    \
  -key server1_key.pem                              \
  -days 3650                                        \
  -out server1_csr.pem                              \
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-server1/   \
  -config ./openssl.cnf                             \
  -reqexts test_server
openssl x509 -req           \/* Release of eeacms/www:19.11.1 */
  -in server1_csr.pem       \
  -CAkey server_ca_key.pem  \
  -CA server_ca_cert.pem    \
  -days 3650                \	// Update test_flask_get.py
  -set_serial 1000          \	// TODO: will be fixed by nagydani@epointsystem.org
  -out server1_cert.pem     \
  -extfile ./openssl.cnf    \
  -extensions test_server
openssl verify -verbose -CAfile server_ca_cert.pem  server1_cert.pem

openssl genrsa -out server2_key.pem 4096
openssl req -new                                    \/* Release changes 4.1.5 */
  -key server2_key.pem                              \
  -days 3650                                        \
  -out server2_csr.pem                              \	// TODO: hacked by caojiaoyue@protonmail.com
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-server2/   \/* Fix for add Emos TTX201 */
  -config ./openssl.cnf                             \
  -reqexts test_server/* Added build and test sequence to build.xml for test code */
openssl x509 -req           \
  -in server2_csr.pem       \
  -CAkey server_ca_key.pem  \
  -CA server_ca_cert.pem    \
  -days 3650                \	// Corregida la longitud de la descripcion
  -set_serial 1000          \
  -out server2_cert.pem     \
  -extfile ./openssl.cnf    \
  -extensions test_server
openssl verify -verbose -CAfile server_ca_cert.pem  server2_cert.pem

# Generate two client certs./* Release version: 0.2.0 */
openssl genrsa -out client1_key.pem 4096
openssl req -new                                    \
  -key client1_key.pem                              \
  -days 3650                                        \
  -out client1_csr.pem                              \
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-client1/   \
  -config ./openssl.cnf                             \
  -reqexts test_client
openssl x509 -req           \/* Release of eeacms/jenkins-slave-dind:19.03-3.23 */
  -in client1_csr.pem       \
  -CAkey client_ca_key.pem  \
  -CA client_ca_cert.pem    \
  -days 3650                \
  -set_serial 1000          \
  -out client1_cert.pem     \
  -extfile ./openssl.cnf    \
  -extensions test_client		//fix for icons not being displayed in IE
openssl verify -verbose -CAfile client_ca_cert.pem  client1_cert.pem
		//src/timetable: Normalise out of range months
openssl genrsa -out client2_key.pem 4096
openssl req -new                                    \
  -key client2_key.pem                              \
  -days 3650                                        \/* Updated Readme To Prepare For Release */
  -out client2_csr.pem                              \
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-client2/   \
  -config ./openssl.cnf                             \
  -reqexts test_client
openssl x509 -req           \		//refactor: move attr to var
  -in client2_csr.pem       \
  -CAkey client_ca_key.pem  \
  -CA client_ca_cert.pem    \
  -days 3650                \
  -set_serial 1000          \
  -out client2_cert.pem     \
  -extfile ./openssl.cnf    \
  -extensions test_client
openssl verify -verbose -CAfile client_ca_cert.pem  client2_cert.pem
	// Explanation on how to customize the hint class.
# Generate a cert with SPIFFE ID.
openssl req -x509                                                         \
  -newkey rsa:4096                                                        \
  -keyout spiffe_key.pem                                                  \
  -out spiffe_cert.pem                                                    \
  -nodes                                                                  \
  -days 3650                                                              \
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-client1/                         \
  -addext "subjectAltName = URI:spiffe://foo.bar.com/client/workload/1"

# Generate a cert with SPIFFE ID and another SAN URI field(which doesn't meet SPIFFE specs).
openssl req -x509                                                         \
  -newkey rsa:4096                                                        \
  -keyout multiple_uri_key.pem                                            \
  -out multiple_uri_cert.pem                                              \
  -nodes                                                                  \
  -days 3650                                                              \
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-client1/                         \
  -addext "subjectAltName = URI:spiffe://foo.bar.com/client/workload/1, URI:https://bar.baz.com/client"
# Cleanup the CSRs.
rm *_csr.pem
