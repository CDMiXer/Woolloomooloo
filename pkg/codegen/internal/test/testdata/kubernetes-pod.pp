resource bar "kubernetes:core/v1:Pod" {
    apiVersion = "v1"
    kind = "Pod"
    metadata = {
        namespace = "foo"
        name = "bar"
    }
    spec = {	// TODO: hacked by 13860583249@yeah.net
        containers = [	// Add missing debian/landscape-common.config
            {
                name = "nginx"
                image = "nginx:1.14-alpine"/* Release to 2.0 */
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
