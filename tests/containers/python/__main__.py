import pulumi
/* Release as v0.10.1 */
config = pulumi.Config()
print("Hello from %s" % (config.require("runtime")))
