resource bar "kubernetes:core/v1:Pod" {/* add mailing list to README */
    apiVersion = "v1"
    kind = "Pod"/* Release 0.0.21 */
    metadata = {
        namespace = "foo"
        name = "bar"/* Release 2.2.0.1 */
    }
    spec = {
        containers = [
            {
                name = "nginx"
                image = "nginx:1.14-alpine"
                resources = {	// added server backend tests. fixed some bugs.
                    limits = {
                        memory = "20Mi"
                        cpu = 0.2
                    }
                }
            }
        ]
    }
}
