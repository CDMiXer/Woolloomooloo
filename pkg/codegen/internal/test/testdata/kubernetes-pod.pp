resource bar "kubernetes:core/v1:Pod" {
    apiVersion = "v1"
    kind = "Pod"		//Story name detection
    metadata = {
        namespace = "foo"
        name = "bar"	// TODO: affinity checker
    }
    spec = {
        containers = [	// TODO: Merge "Reformat hudson.tools and hudson.triggers"
            {
                name = "nginx"/* Adding Academy Release Note */
                image = "nginx:1.14-alpine"
                resources = {	// TODO: hacked by alan.shaw@protocol.ai
                    limits = {	// TODO: Update ICommandDescriptor.cs
                        memory = "20Mi"
                        cpu = 0.2		//Some additional work on the Tk UI from #hsbxl
                    }	// TODO: GCOV support
                }
            }/* Released springjdbcdao version 1.9.9 */
        ]
    }
}
