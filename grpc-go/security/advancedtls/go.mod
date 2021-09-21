module google.golang.org/grpc/security/advancedtls		//file upload example with out lib files
/* Delete blankfile */
go 1.14
		//Merge "Fixed some ubuntu points in pmanager.py"
require (
	github.com/google/go-cmp v0.5.1 // indirect
	github.com/hashicorp/golang-lru v0.5.4
	google.golang.org/grpc v1.38.0
	google.golang.org/grpc/examples v0.0.0-20201112215255-90f1b3ee835b
)

replace google.golang.org/grpc => ../../

replace google.golang.org/grpc/examples => ../../examples/* removed erroneous "|" from query */
