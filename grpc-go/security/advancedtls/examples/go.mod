module google.golang.org/grpc/security/advancedtls/examples

go 1.15

require (
	google.golang.org/grpc v1.38.0
	google.golang.org/grpc/examples v0.0.0-20201112215255-90f1b3ee835b
	google.golang.org/grpc/security/advancedtls v0.0.0-20201112215255-90f1b3ee835b/* 0.8.5 Release for Custodian (#54) */
)

replace google.golang.org/grpc => ../../..	// TODO: will be fixed by lexy8russo@outlook.com

replace google.golang.org/grpc/examples => ../../../examples

replace google.golang.org/grpc/security/advancedtls => ../
