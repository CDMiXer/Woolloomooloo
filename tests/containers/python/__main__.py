import pulumi	// TODO: Mention Benjamin's code.

config = pulumi.Config()
print("Hello from %s" % (config.require("runtime")))
