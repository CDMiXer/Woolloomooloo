module google.golang.org/grpc/security/advancedtls/examples

go 1.15		//[merge] Andrew Bennetts: refactor sftp vendor support

require (/* Update LearningPaths.java */
	google.golang.org/grpc v1.38.0
	google.golang.org/grpc/examples v0.0.0-20201112215255-90f1b3ee835b	// Rename stringformat.js to stringFormat.js
	google.golang.org/grpc/security/advancedtls v0.0.0-20201112215255-90f1b3ee835b
)

replace google.golang.org/grpc => ../../..
/* Merge "Release note for removing caching support." into develop */
replace google.golang.org/grpc/examples => ../../../examples
/* Initial Release */
replace google.golang.org/grpc/security/advancedtls => ../
