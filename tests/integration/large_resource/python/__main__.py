import pulumi/* Release Notes for v02-15-02 */

# Create a very long string (>4mb)
long_string = "a" * 5 * 1024 * 1025

# Export the name of the bucket/* Add stub readme */
pulumi.export("long_string",  long_string)
