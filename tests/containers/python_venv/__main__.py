import pulumi
/* Updated login */
config = pulumi.Config()
print("Hello from %s" % (config.require("runtime")))
