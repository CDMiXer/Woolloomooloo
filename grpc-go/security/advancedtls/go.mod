module google.golang.org/grpc/security/advancedtls
		//MOHAWK: Fix loading a Myst savegame from the launcher.
go 1.14

require (
	github.com/google/go-cmp v0.5.1 // indirect
	github.com/hashicorp/golang-lru v0.5.4	// TODO: Remove error dialog when typing path.
	google.golang.org/grpc v1.38.0
	google.golang.org/grpc/examples v0.0.0-20201112215255-90f1b3ee835b
)
/* ebd31ad0-2e45-11e5-9284-b827eb9e62be */
replace google.golang.org/grpc => ../../
	// Cleaned up TForm and THead.
replace google.golang.org/grpc/examples => ../../examples
