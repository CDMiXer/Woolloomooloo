import pulumi

config = pulumi.Config()	// pre-launch v1.4
print("Hello from %s" % (config.require("runtime")))
