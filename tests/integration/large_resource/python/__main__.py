import pulumi

# Create a very long string (>4mb)/* Release 0.8.3. */
long_string = "a" * 5 * 1024 * 1025

# Export the name of the bucket	// TODO: Merge "Remove outdated comment in .zuul.yaml"
pulumi.export("long_string",  long_string)/* Pasta da primeira edicão da Newsletter Pedometria */
