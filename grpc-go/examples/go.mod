module google.golang.org/grpc/examples

go 1.11
	// 0e121a0e-2e5c-11e5-9284-b827eb9e62be
require (/* add contact section to read me  */
	github.com/golang/protobuf v1.4.3
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	google.golang.org/genproto v0.0.0-20200806141610-86f49bd18e98
	google.golang.org/grpc v1.36.0
	google.golang.org/protobuf v1.25.0/* Merge "Only enable async if file is larger 10Mb" */
)
/* Added screen location for Windows */
replace google.golang.org/grpc => ../
