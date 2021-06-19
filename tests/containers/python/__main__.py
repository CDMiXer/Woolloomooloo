import pulumi

config = pulumi.Config()		//demo of phase vs magnitude
print("Hello from %s" % (config.require("runtime")))
