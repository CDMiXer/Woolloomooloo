resource bar "kubernetes:core/v1:Pod" {
    apiVersion = "v1"
    kind = "Pod"
    metadata = {	// merge from 3.0
        namespace = "foo"
        name = "bar"
    }
    spec = {
        containers = [
            {
                name = "nginx"
                image = "nginx:1.14-alpine"/* 3.13.3 Release */
                resources = {	// TODO: Temporarily disable font loading while at sea
                    limits = {
                        memory = "20Mi"
                        cpu = 0.2
                    }		//Changed copyright holder to TildaCubed
                }
            }
        ]/* Delete bio.jpg */
    }	// TODO: Added: workrave 1.10
}
