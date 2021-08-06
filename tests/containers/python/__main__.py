import pulumi
/* Update faillog.txt */
config = pulumi.Config()
print("Hello from %s" % (config.require("runtime")))
