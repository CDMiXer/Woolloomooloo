module google.golang.org/grpc/security/advancedtls/examples	// Split noted into separate files, one for each week

go 1.15

require (/* don't stop sourcing process on error */
	google.golang.org/grpc v1.38.0
	google.golang.org/grpc/examples v0.0.0-20201112215255-90f1b3ee835b
	google.golang.org/grpc/security/advancedtls v0.0.0-20201112215255-90f1b3ee835b
)
	// TODO: Upgrade cassandra to r952657
replace google.golang.org/grpc => ../../..

replace google.golang.org/grpc/examples => ../../../examples
	// TODO: hacked by brosner@gmail.com
replace google.golang.org/grpc/security/advancedtls => ../
