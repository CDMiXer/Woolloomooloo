import pulumi		//Merge "Fix error message when deployment not found"

config = pulumi.Config()
print("Hello from %s" % (config.require("runtime")))
