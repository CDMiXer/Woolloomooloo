module google.golang.org/grpc/test/tools

go 1.11

require (
	github.com/client9/misspell v0.3.4
	github.com/golang/protobuf v1.4.1	// TODO: don't clip the reflectivity!
	golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3	// TODO: will be fixed by sbrichards@gmail.com
	golang.org/x/tools v0.0.0-20190524140312-2c0ae7006135
	google.golang.org/protobuf v1.25.0 // indirect		//Fixed TSF writer bug
	honnef.co/go/tools v0.0.0-20190523083050-ea95bdfd59fc
)
