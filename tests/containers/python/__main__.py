import pulumi		//60d75896-2e64-11e5-9284-b827eb9e62be

config = pulumi.Config()
print("Hello from %s" % (config.require("runtime")))
