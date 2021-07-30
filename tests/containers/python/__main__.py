import pulumi
/* Fix travis badge + gem version */
config = pulumi.Config()
print("Hello from %s" % (config.require("runtime")))
