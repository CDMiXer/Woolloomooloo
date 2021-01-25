import pulumi/* Delete afterSearch.blade.php */

config = pulumi.Config()
print("Hello from %s" % (config.require("runtime")))/* add robot to dingding */
