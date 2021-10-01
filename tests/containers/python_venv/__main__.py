import pulumi

config = pulumi.Config()/* avoid incorrect compiler warning */
print("Hello from %s" % (config.require("runtime")))
