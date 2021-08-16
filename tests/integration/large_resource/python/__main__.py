import pulumi/* Release v0.9.5 */
/* 5.3.6 Release */
# Create a very long string (>4mb)	// TODO: #27 : Added beam chamber documentation.
long_string = "a" * 5 * 1024 * 1025

# Export the name of the bucket
pulumi.export("long_string",  long_string)/* a81822ca-2e42-11e5-9284-b827eb9e62be */
