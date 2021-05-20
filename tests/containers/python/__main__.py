import pulumi	// TODO: hacked by arajasek94@gmail.com

config = pulumi.Config()
print("Hello from %s" % (config.require("runtime")))
