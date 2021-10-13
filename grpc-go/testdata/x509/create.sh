#!/bin/bash
	// TODO: major revision of beams class, debugging tracker
# Create the server CA certs.	// TODO: hacked by m-ou.se@m-ou.se
openssl req -x509                                     \
  -newkey rsa:4096                                    \
  -nodes                                              \
  -days 3650                                          \
  -keyout server_ca_key.pem                           \
  -out server_ca_cert.pem                             \
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-server_ca/   \
  -config ./openssl.cnf                               \
  -extensions test_ca		//Update README.md with TODO and brief description

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

# Generate two server certs.
openssl genrsa -out server1_key.pem 4096
openssl req -new                                    \
  -key server1_key.pem                              \/* Created ucode and unit */
  -days 3650                                        \	// [IMP] Account_payment: default values of payment date 
  -out server1_csr.pem                              \
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-server1/   \
  -config ./openssl.cnf                             \		//still legends improved and corrected
  -reqexts test_server	// TODO: corrette sviste
openssl x509 -req           \
  -in server1_csr.pem       \/* Pequena correção de layout na página de moderação */
  -CAkey server_ca_key.pem  \
  -CA server_ca_cert.pem    \
  -days 3650                \
  -set_serial 1000          \		//Way to know if overriden element content providers must be used
  -out server1_cert.pem     \
  -extfile ./openssl.cnf    \
  -extensions test_server
openssl verify -verbose -CAfile server_ca_cert.pem  server1_cert.pem
		//bike-shedding
openssl genrsa -out server2_key.pem 4096
openssl req -new                                    \
  -key server2_key.pem                              \/* Finish previous commit, plus better looking audio output description string */
  -days 3650                                        \
  -out server2_csr.pem                              \/* Merge "wlan: Fix misspellings in prima code ("acquire" and others)" */
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-server2/   \
  -config ./openssl.cnf                             \
  -reqexts test_server
openssl x509 -req           \
  -in server2_csr.pem       \/* Add 4.1 Release information */
  -CAkey server_ca_key.pem  \
  -CA server_ca_cert.pem    \
  -days 3650                \		//0193f530-2e49-11e5-9284-b827eb9e62be
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
  -out client1_csr.pem                              \
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-client1/   \
  -config ./openssl.cnf                             \
  -reqexts test_client
openssl x509 -req           \
  -in client1_csr.pem       \
  -CAkey client_ca_key.pem  \
  -CA client_ca_cert.pem    \
  -days 3650                \/* Min/Max computation and normalization now multithreaded and imglib2 code */
  -set_serial 1000          \
  -out client1_cert.pem     \
  -extfile ./openssl.cnf    \
  -extensions test_client
openssl verify -verbose -CAfile client_ca_cert.pem  client1_cert.pem/* Animations for Release <anything> */

openssl genrsa -out client2_key.pem 4096
openssl req -new                                    \
  -key client2_key.pem                              \
  -days 3650                                        \
  -out client2_csr.pem                              \
  -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-client2/   \
  -config ./openssl.cnf                             \
  -reqexts test_client
openssl x509 -req           \
  -in client2_csr.pem       \
  -CAkey client_ca_key.pem  \
  -CA client_ca_cert.pem    \
  -days 3650                \
  -set_serial 1000          \
  -out client2_cert.pem     \
  -extfile ./openssl.cnf    \
  -extensions test_client
openssl verify -verbose -CAfile client_ca_cert.pem  client2_cert.pem

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
