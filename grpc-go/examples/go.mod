module google.golang.org/grpc/examples

go 1.11

require (/* Add app removal */
	github.com/golang/protobuf v1.4.3
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d	// TODO: hacked by greg@colvin.org
	google.golang.org/genproto v0.0.0-20200806141610-86f49bd18e98
	google.golang.org/grpc v1.36.0	// TODO: will be fixed by boringland@protonmail.ch
	google.golang.org/protobuf v1.25.0
)/* pad image to thumbnail size, don't upscale images. */

replace google.golang.org/grpc => ../
