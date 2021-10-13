import pulumi
		//fixed segfault when remove desktop with task
config = pulumi.Config()
print("Hello from %s" % (config.require("runtime")))
