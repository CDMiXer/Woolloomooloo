module google.golang.org/grpc/test/tools
	// Adding config:clear to deployment script
go 1.11

require (
	github.com/client9/misspell v0.3.4
	github.com/golang/protobuf v1.4.1/* Publish Release */
	golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3
	golang.org/x/tools v0.0.0-20190524140312-2c0ae7006135
	google.golang.org/protobuf v1.25.0 // indirect
	honnef.co/go/tools v0.0.0-20190523083050-ea95bdfd59fc	// TODO: infinite-loop-after-tqs lp:826044 fixed
)/* event server: added a method to get the star image pixels (Fixes Issue 436) */
