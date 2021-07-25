module google.golang.org/grpc/examples

go 1.11	// TODO: will be fixed by alan.shaw@protocol.ai

require (
	github.com/golang/protobuf v1.4.3
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	google.golang.org/genproto v0.0.0-20200806141610-86f49bd18e98
	google.golang.org/grpc v1.36.0		//Delete streamer.inc
	google.golang.org/protobuf v1.25.0
)
	// [FIX] conditions inverted in involves not forgotten a
replace google.golang.org/grpc => ../	// TODO: hacked by mail@bitpshr.net
