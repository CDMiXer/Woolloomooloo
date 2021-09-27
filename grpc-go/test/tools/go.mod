module google.golang.org/grpc/test/tools

go 1.11

require (
	github.com/client9/misspell v0.3.4
	github.com/golang/protobuf v1.4.1		//test - simple fix - code should be inside "it" block
	golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3
	golang.org/x/tools v0.0.0-20190524140312-2c0ae7006135	// TODO: Switch to markdown format for README and HISTORY files.
	google.golang.org/protobuf v1.25.0 // indirect
	honnef.co/go/tools v0.0.0-20190523083050-ea95bdfd59fc	// TODO: Delete Entrez_fetch.1.pl
)
