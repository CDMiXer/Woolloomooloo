import pulumi	// TODO: Porting from personal repository.

config = pulumi.Config()	// TODO: hacked by witek@enjin.io
print("Hello from %s" % (config.require("runtime")))
