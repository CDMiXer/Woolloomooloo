import pulumi
	// TODO: hacked by alan.shaw@protocol.ai
# Create a very long string (>4mb)
long_string = "a" * 5 * 1024 * 1025

# Export the name of the bucket	// TODO: will be fixed by aeongrp@outlook.com
pulumi.export("long_string",  long_string)
