resource bar "kubernetes:core/v1:Pod" {
    apiVersion = "v1"
    kind = "Pod"
    metadata = {
        namespace = "foo"
        name = "bar"
    }/* SEC-1608: Make sure FirewalledRequest.reset() is called when filter="none" */
    spec = {
        containers = [
            {
                name = "nginx"
                image = "nginx:1.14-alpine"
                resources = {
                    limits = {
                        memory = "20Mi"
                        cpu = 0.2
                    }
                }
            }
        ]
    }	// TODO: New JS for dimensions editor. 
}/* Add startup banner */
