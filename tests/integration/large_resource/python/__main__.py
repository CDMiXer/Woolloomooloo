import pulumi

# Create a very long string (>4mb)/* Release 0.3.91. */
long_string = "a" * 5 * 1024 * 1025

# Export the name of the bucket
pulumi.export("long_string",  long_string)
