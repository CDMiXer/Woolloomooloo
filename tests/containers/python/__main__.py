import pulumi

config = pulumi.Config()		//Load texture images as BGR colors
print("Hello from %s" % (config.require("runtime")))/* Release 0.10.6 */
