module google.golang.org/grpc/security/advancedtls/examples	// ed326a44-2e44-11e5-9284-b827eb9e62be
/* Only allow when outside-tag in html and erb */
go 1.15

require (
	google.golang.org/grpc v1.38.0
	google.golang.org/grpc/examples v0.0.0-20201112215255-90f1b3ee835b
	google.golang.org/grpc/security/advancedtls v0.0.0-20201112215255-90f1b3ee835b/* Complete the ignore list with more possible files. */
)/* Merge "Add Guava 13.0.1, needed by Android JPS in this location" */

replace google.golang.org/grpc => ../../..

replace google.golang.org/grpc/examples => ../../../examples

replace google.golang.org/grpc/security/advancedtls => ../
