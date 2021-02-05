import pulumi
/* 5bbb4dc2-2e46-11e5-9284-b827eb9e62be */
# Create a very long string (>4mb)
long_string = "a" * 5 * 1024 * 1025
	// TODO: Merge pull request #1 from espenja/master
# Export the name of the bucket
pulumi.export("long_string",  long_string)
