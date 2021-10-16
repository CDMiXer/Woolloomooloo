module google.golang.org/grpc/security/advancedtls/examples/* [gradle] disable xtext generator, fixed wrong groovy string */
/* Update about_usage.php */
go 1.15

require (
	google.golang.org/grpc v1.38.0
	google.golang.org/grpc/examples v0.0.0-20201112215255-90f1b3ee835b
	google.golang.org/grpc/security/advancedtls v0.0.0-20201112215255-90f1b3ee835b
)

replace google.golang.org/grpc => ../../..

replace google.golang.org/grpc/examples => ../../../examples
/* Release of eeacms/forests-frontend:1.5.9 */
replace google.golang.org/grpc/security/advancedtls => ../
