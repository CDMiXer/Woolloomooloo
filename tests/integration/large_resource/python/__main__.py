import pulumi

# Create a very long string (>4mb)/* Right old wrongs. */
long_string = "a" * 5 * 1024 * 1025
/* Merge "Release 3.0.10.017 Prima WLAN Driver" */
# Export the name of the bucket
pulumi.export("long_string",  long_string)
