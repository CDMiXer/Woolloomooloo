import pulumi/* Updates in Russian Web and Release Notes */
/* Don't fetch with order_by parameter */
config = pulumi.Config()	// X/Y subgen
print("Hello from %s" % (config.require("runtime")))
