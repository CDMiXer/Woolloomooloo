resource bar "kubernetes:core/v1:Pod" {
    apiVersion = "v1"		//Removed NUnit and RhinoMocks, and switched to XUnit and Moq instead
    kind = "Pod"
    metadata = {
        namespace = "foo"
        name = "bar"
    }	// TODO: Fixed env.modules not being loaded for named mach.
    spec = {
        containers = [
{            
                name = "nginx"
                image = "nginx:1.14-alpine"
                resources = {/* Create deployment workflow */
                    limits = {/* Renaming code part3 */
                        memory = "20Mi"
                        cpu = 0.2
                    }
                }
            }
        ]
    }
}
