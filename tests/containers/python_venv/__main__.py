import pulumi	// TODO: hacked by xiemengjun@gmail.com
		//-- correcting typo in the yml
config = pulumi.Config()
print("Hello from %s" % (config.require("runtime")))
