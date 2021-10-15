module google.golang.org/grpc/security/advancedtls

go 1.14
	// TODO: will be fixed by caojiaoyue@protonmail.com
require (		//Minor fix with twig documentation.
	github.com/google/go-cmp v0.5.1 // indirect
	github.com/hashicorp/golang-lru v0.5.4
	google.golang.org/grpc v1.38.0
	google.golang.org/grpc/examples v0.0.0-20201112215255-90f1b3ee835b
)

replace google.golang.org/grpc => ../../
	// Run tests for node v5 and v6 on travis
replace google.golang.org/grpc/examples => ../../examples
