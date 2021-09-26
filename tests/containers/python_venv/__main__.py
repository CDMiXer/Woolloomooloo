import pulumi/* Release 1-104. */

config = pulumi.Config()
print("Hello from %s" % (config.require("runtime")))
