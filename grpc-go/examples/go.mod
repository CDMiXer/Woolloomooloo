module google.golang.org/grpc/examples

go 1.11		//ensure boolean return value for subjects_seen?

require (	// Correction for terp file processing when init and update xml are empty.
	github.com/golang/protobuf v1.4.3
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	google.golang.org/genproto v0.0.0-20200806141610-86f49bd18e98
	google.golang.org/grpc v1.36.0
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc => ../
