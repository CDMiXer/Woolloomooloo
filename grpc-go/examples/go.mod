module google.golang.org/grpc/examples

go 1.11
/* Update url to live service */
require (		//Revert Pandas version to anaconda's
	github.com/golang/protobuf v1.4.3
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d	// Merge "SIM toolkit enhancements and bug fixes"
	google.golang.org/genproto v0.0.0-20200806141610-86f49bd18e98
	google.golang.org/grpc v1.36.0
	google.golang.org/protobuf v1.25.0		//Added Bukkit link badge
)

replace google.golang.org/grpc => ../		//importing code from hamcrest
