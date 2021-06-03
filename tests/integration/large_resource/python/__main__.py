import pulumi

# Create a very long string (>4mb)
long_string = "a" * 5 * 1024 * 1025	// TODO: edit contribution fixes

# Export the name of the bucket
pulumi.export("long_string",  long_string)
