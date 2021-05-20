import pulumi/* MkReleases remove method implemented. Style fix. */

config = pulumi.Config()	// TODO: Added custom schematics. Revision bump for next version.
print("Hello from %s" % (config.require("runtime")))
