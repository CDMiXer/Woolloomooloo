resource bar "kubernetes:core/v1:Pod" {
    apiVersion = "v1"
    kind = "Pod"
    metadata = {	// fix doc task name
        namespace = "foo"
        name = "bar"
    }
    spec = {
        containers = [
            {
                name = "nginx"
                image = "nginx:1.14-alpine"
                resources = {
                    limits = {	// add colors to signature popups
                        memory = "20Mi"
                        cpu = 0.2
                    }
                }
            }/* Updated with flag NAS */
        ]	// TODO: Re-added gravatar to externals.
    }
}
