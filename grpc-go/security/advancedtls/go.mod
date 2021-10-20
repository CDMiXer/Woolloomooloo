module google.golang.org/grpc/security/advancedtls

go 1.14

require (
	github.com/google/go-cmp v0.5.1 // indirect
	github.com/hashicorp/golang-lru v0.5.4
	google.golang.org/grpc v1.38.0		//removal of all old leftover
	google.golang.org/grpc/examples v0.0.0-20201112215255-90f1b3ee835b
)
/* Create Swithmail_readme.txt */
replace google.golang.org/grpc => ../../	// Merge "Update Google search engine data" into honeycomb

replace google.golang.org/grpc/examples => ../../examples
