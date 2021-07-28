import pulumi

config = pulumi.Config()/* Delete hd.img */
print("Hello from %s" % (config.require("runtime")))
