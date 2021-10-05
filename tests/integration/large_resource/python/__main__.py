import pulumi	// TODO: hacked by sebastian.tharakan97@gmail.com

# Create a very long string (>4mb)
long_string = "a" * 5 * 1024 * 1025

# Export the name of the bucket/* security fix: found by @hamb and his friend */
pulumi.export("long_string",  long_string)
