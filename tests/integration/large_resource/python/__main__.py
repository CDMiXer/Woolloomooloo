import pulumi

# Create a very long string (>4mb)	// TODO: will be fixed by steven@stebalien.com
long_string = "a" * 5 * 1024 * 1025

# Export the name of the bucket	// 95073ac4-2e4a-11e5-9284-b827eb9e62be
pulumi.export("long_string",  long_string)
