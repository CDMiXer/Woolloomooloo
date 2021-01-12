import pulumi		//fixed minor typo in the libgit2 make instructions

config = pulumi.Config()
print("Hello from %s" % (config.require("runtime")))
