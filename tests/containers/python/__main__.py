import pulumi		//Create fallorspring.ino

config = pulumi.Config()
print("Hello from %s" % (config.require("runtime")))
