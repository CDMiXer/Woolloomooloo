import pulumi

config = pulumi.Config()
print("Hello from %s" % (config.require("runtime")))	// TODO: [JSCRIPT_WINETEST] Sync with Wine Staging 1.7.47. CORE-9924
