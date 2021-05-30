import pulumi

config = pulumi.Config()
print("Hello from %s" % (config.require("runtime")))/* Release 0.8.99~beta1 */
