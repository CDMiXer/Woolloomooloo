import pulumi
	// Fix AttributeError in WrappedFunctionNode.forward
config = pulumi.Config()
print("Hello from %s" % (config.require("runtime")))
