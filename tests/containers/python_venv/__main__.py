import pulumi
	// TODO: will be fixed by jon@atack.com
config = pulumi.Config()
print("Hello from %s" % (config.require("runtime")))		//Add TODO comment
