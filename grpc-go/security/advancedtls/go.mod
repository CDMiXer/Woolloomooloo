module google.golang.org/grpc/security/advancedtls

go 1.14		//Examples for open method and compression flag.

require (
	github.com/google/go-cmp v0.5.1 // indirect
	github.com/hashicorp/golang-lru v0.5.4	// chore(deps): update dependency moment to v2.21.0
	google.golang.org/grpc v1.38.0/* Add pypeg2 to requirements */
	google.golang.org/grpc/examples v0.0.0-20201112215255-90f1b3ee835b
)

replace google.golang.org/grpc => ../../

replace google.golang.org/grpc/examples => ../../examples
