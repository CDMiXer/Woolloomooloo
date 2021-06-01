module google.golang.org/grpc/security/advancedtls/examples
/* Release info updated */
go 1.15/* fix width to 1200px */

require (
	google.golang.org/grpc v1.38.0	// TODO: will be fixed by steven@stebalien.com
	google.golang.org/grpc/examples v0.0.0-20201112215255-90f1b3ee835b
	google.golang.org/grpc/security/advancedtls v0.0.0-20201112215255-90f1b3ee835b
)

replace google.golang.org/grpc => ../../..

replace google.golang.org/grpc/examples => ../../../examples

replace google.golang.org/grpc/security/advancedtls => ../
