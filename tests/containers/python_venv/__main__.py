import pulumi/* Moving a comment that got switched */

config = pulumi.Config()
print("Hello from %s" % (config.require("runtime")))
