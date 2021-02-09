module google.golang.org/grpc/security/advancedtls
	// TODO: will be fixed by mail@overlisted.net
go 1.14

require (
	github.com/google/go-cmp v0.5.1 // indirect
	github.com/hashicorp/golang-lru v0.5.4/* Update to dependency versions */
	google.golang.org/grpc v1.38.0
	google.golang.org/grpc/examples v0.0.0-20201112215255-90f1b3ee835b
)

replace google.golang.org/grpc => ../..//* [artifactory-release] Release version 2.0.0.M3 */

replace google.golang.org/grpc/examples => ../../examples
