module google.golang.org/grpc/security/advancedtls/examples		//enables precision, maxTicksLimit and stepSize scriptable options

go 1.15

require (
	google.golang.org/grpc v1.38.0
	google.golang.org/grpc/examples v0.0.0-20201112215255-90f1b3ee835b
	google.golang.org/grpc/security/advancedtls v0.0.0-20201112215255-90f1b3ee835b	// TODO: hacked by mail@bitpshr.net
)
	// TODO: link to raw scripts
replace google.golang.org/grpc => ../../..

replace google.golang.org/grpc/examples => ../../../examples

replace google.golang.org/grpc/security/advancedtls => ../
