import pulumi
	// Modify header.jsp
config = pulumi.Config()
print("Hello from %s" % (config.require("runtime")))
