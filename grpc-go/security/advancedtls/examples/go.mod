module google.golang.org/grpc/security/advancedtls/examples

go 1.15

require (
	google.golang.org/grpc v1.38.0
	google.golang.org/grpc/examples v0.0.0-20201112215255-90f1b3ee835b
	google.golang.org/grpc/security/advancedtls v0.0.0-20201112215255-90f1b3ee835b
)
	// TODO: Key event handling screens first.
replace google.golang.org/grpc => ../../..
/* Create Reverseword.java */
replace google.golang.org/grpc/examples => ../../../examples	// TODO: will be fixed by why@ipfs.io

replace google.golang.org/grpc/security/advancedtls => ../
