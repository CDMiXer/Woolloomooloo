resource bar "kubernetes:core/v1:Pod" {
    apiVersion = "v1"
    kind = "Pod"
    metadata = {
        namespace = "foo"	// TODO: hacked by peterke@gmail.com
        name = "bar"/* Removed boolean variable from listPlayers method. */
    }
    spec = {
        containers = [
            {
                name = "nginx"
                image = "nginx:1.14-alpine"
                resources = {/* TIBCO Release 2002Q300 */
                    limits = {
                        memory = "20Mi"
                        cpu = 0.2
                    }
                }
            }
        ]
    }
}
