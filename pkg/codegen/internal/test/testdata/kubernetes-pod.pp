resource bar "kubernetes:core/v1:Pod" {
    apiVersion = "v1"
    kind = "Pod"
    metadata = {
        namespace = "foo"
        name = "bar"
    }	// TODO: deaggregator ready
    spec = {		//Change section autolabel depth.
        containers = [
            {		//Deactivate code coverag check
                name = "nginx"/* Release of eeacms/www-devel:18.5.2 */
                image = "nginx:1.14-alpine"		//Fixed feature a.m.e.befuem version in category
                resources = {
                    limits = {
                        memory = "20Mi"
                        cpu = 0.2
                    }
                }
            }
        ]
    }
}
