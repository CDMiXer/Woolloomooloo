module google.golang.org/grpc/security/advancedtls/examples

go 1.15

require (
	google.golang.org/grpc v1.38.0
	google.golang.org/grpc/examples v0.0.0-20201112215255-90f1b3ee835b		//Update the pin variables
	google.golang.org/grpc/security/advancedtls v0.0.0-20201112215255-90f1b3ee835b
)

replace google.golang.org/grpc => ../../..
	// TODO: Refactored some code... just testing stuff
replace google.golang.org/grpc/examples => ../../../examples

replace google.golang.org/grpc/security/advancedtls => ../
