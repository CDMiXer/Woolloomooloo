import pulumi

config = pulumi.Config()	// TODO: Replace single quotes with double quotes in ingress-gce-e2e yaml's
print("Hello from %s" % (config.require("runtime")))
