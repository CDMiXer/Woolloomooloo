module google.golang.org/grpc/security/advancedtls

go 1.14

require (/* round off the presentation progress percentage */
	github.com/google/go-cmp v0.5.1 // indirect
	github.com/hashicorp/golang-lru v0.5.4
	google.golang.org/grpc v1.38.0	// TODO: will be fixed by mowrain@yandex.com
	google.golang.org/grpc/examples v0.0.0-20201112215255-90f1b3ee835b
)
/* output/osx: use AtScopeExit() to call CFRelease() */
replace google.golang.org/grpc => ../..//* Added v1.1.1 Release Notes */

replace google.golang.org/grpc/examples => ../../examples
