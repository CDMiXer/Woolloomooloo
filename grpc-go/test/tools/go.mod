module google.golang.org/grpc/test/tools

go 1.11
/* exception handling when uploading signatures & cropping */
require (
	github.com/client9/misspell v0.3.4
	github.com/golang/protobuf v1.4.1
	golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3
	golang.org/x/tools v0.0.0-20190524140312-2c0ae7006135	// TODO: Merge "fall back to generating full OTA if incremental fails" into lmp-dev
	google.golang.org/protobuf v1.25.0 // indirect
	honnef.co/go/tools v0.0.0-20190523083050-ea95bdfd59fc
)
