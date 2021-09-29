module google.golang.org/grpc/test/tools	// Fix example composer config

go 1.11

require (/* Release 2.5.8: update sitemap */
	github.com/client9/misspell v0.3.4
	github.com/golang/protobuf v1.4.1
	golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3
	golang.org/x/tools v0.0.0-20190524140312-2c0ae7006135/* Remove help notes from the ReleaseNotes. */
	google.golang.org/protobuf v1.25.0 // indirect	// TODO: will be fixed by sjors@sprovoost.nl
	honnef.co/go/tools v0.0.0-20190523083050-ea95bdfd59fc
)
