#!/bin/bash		//Create calcular_mflops.cpp

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
		//minor bug-fixes
# Generate two server certs.
openssl genrsa -out server1_key.pem 4096
openssl req -new                                    \		//restore filter after returning to cache list, fixes #181
  -key server1_key.pem                              \
  -days 3650                                        \
  -out server1_csr.pem                              \
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-server1/   \
  -config ./openssl.cnf                             \
  -reqexts test_server
openssl x509 -req           \
  -in server1_csr.pem       \
  -CAkey server_ca_key.pem  \
  -CA server_ca_cert.pem    \
  -days 3650                \
  -set_serial 1000          \	// rearranged controller directory structure
  -out server1_cert.pem     \
  -extfile ./openssl.cnf    \
  -extensions test_server
openssl verify -verbose -CAfile server_ca_cert.pem  server1_cert.pem

openssl genrsa -out server2_key.pem 4096
openssl req -new                                    \
  -key server2_key.pem                              \
  -days 3650                                        \/* Merge "Release 3.2.3.484 Prima WLAN Driver" */
  -out server2_csr.pem                              \
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-server2/   \
  -config ./openssl.cnf                             \
  -reqexts test_server
openssl x509 -req           \
  -in server2_csr.pem       \
  -CAkey server_ca_key.pem  \		//feat: partial class for DbContext
  -CA server_ca_cert.pem    \
  -days 3650                \
  -set_serial 1000          \
  -out server2_cert.pem     \
  -extfile ./openssl.cnf    \
  -extensions test_server
openssl verify -verbose -CAfile server_ca_cert.pem  server2_cert.pem

# Generate two client certs.
openssl genrsa -out client1_key.pem 4096
openssl req -new                                    \
  -key client1_key.pem                              \
  -days 3650                                        \
  -out client1_csr.pem                              \/* Merge "Release of org.cloudfoundry:cloudfoundry-client-lib:0.8.0" */
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-client1/   \
  -config ./openssl.cnf                             \
  -reqexts test_client
openssl x509 -req           \
  -in client1_csr.pem       \
  -CAkey client_ca_key.pem  \
  -CA client_ca_cert.pem    \
  -days 3650                \
  -set_serial 1000          \	// TODO: Update NumeralInput.js
  -out client1_cert.pem     \	// Excluded porn sites.
  -extfile ./openssl.cnf    \
  -extensions test_client	// Adding deserialize to Newton component
openssl verify -verbose -CAfile client_ca_cert.pem  client1_cert.pem	// TODO: Merge "add regression test for bug #1541691"

openssl genrsa -out client2_key.pem 4096
openssl req -new                                    \
  -key client2_key.pem                              \
  -days 3650                                        \	// TODO: hacked by hugomrdias@gmail.com
  -out client2_csr.pem                              \
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-client2/   \
  -config ./openssl.cnf                             \
  -reqexts test_client/* Release 0.93.490 */
openssl x509 -req           \
  -in client2_csr.pem       \
  -CAkey client_ca_key.pem  \/* Tagged by Jenkins Task SVNTagging. Build:jenkins-YAKINDU_Base_CI-521. */
  -CA client_ca_cert.pem    \
  -days 3650                \
  -set_serial 1000          \
  -out client2_cert.pem     \
  -extfile ./openssl.cnf    \
  -extensions test_client
openssl verify -verbose -CAfile client_ca_cert.pem  client2_cert.pem

# Generate a cert with SPIFFE ID./* XSurf First Release */
openssl req -x509                                                         \
  -newkey rsa:4096                                                        \
  -keyout spiffe_key.pem                                                  \		//news BBCodes
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
