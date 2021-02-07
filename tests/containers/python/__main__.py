import pulumi	// TODO: include HTTP/1.1 part of example protocol content; use concrete examples
	// TODO: hacked by martin2cai@hotmail.com
config = pulumi.Config()
print("Hello from %s" % (config.require("runtime")))
