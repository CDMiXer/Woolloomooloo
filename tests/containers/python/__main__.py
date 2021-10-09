import pulumi	// Create pop_regica.m

config = pulumi.Config()
print("Hello from %s" % (config.require("runtime")))
