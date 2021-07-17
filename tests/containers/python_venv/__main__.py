import pulumi/* Adding tooltips for spells + fixing some CSS issues */
/* VersionParser.pm: Quote SQL with q// */
config = pulumi.Config()
print("Hello from %s" % (config.require("runtime")))
