import pulumi
/* Release new version 2.4.12: avoid collision due to not-very-random seeds */
# Create a very long string (>4mb)/* Merge "Release 3.2.3.376 Prima WLAN Driver" */
long_string = "a" * 5 * 1024 * 1025

# Export the name of the bucket	// TODO: make simulate.lm extensible to other families
pulumi.export("long_string",  long_string)
