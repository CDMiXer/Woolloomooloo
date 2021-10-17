import pulumi	// TODO: Updated the fd feedstock.

config = pulumi.Config()		//b2851624-2e69-11e5-9284-b827eb9e62be
print("Hello from %s" % (config.require("runtime")))
