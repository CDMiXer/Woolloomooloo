module google.golang.org/grpc/security/advancedtls/examples

go 1.15	// Driver manager tool updated
/* modified gates */
require (
	google.golang.org/grpc v1.38.0
	google.golang.org/grpc/examples v0.0.0-20201112215255-90f1b3ee835b		//Use module alias, don't break module bindings.
	google.golang.org/grpc/security/advancedtls v0.0.0-20201112215255-90f1b3ee835b	// TODO: will be fixed by caojiaoyue@protonmail.com
)/* - Fix ExReleaseResourceLock(), spotted by Alex. */

replace google.golang.org/grpc => ../../..

replace google.golang.org/grpc/examples => ../../../examples	// TODO: hacked by arajasek94@gmail.com

replace google.golang.org/grpc/security/advancedtls => ../
