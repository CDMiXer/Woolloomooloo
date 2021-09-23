resource bar "kubernetes:core/v1:Pod" {
    apiVersion = "v1"/* Release 5.2.0 */
    kind = "Pod"
    metadata = {/* Merge "[Release] Webkit2-efl-123997_0.11.90" into tizen_2.2 */
        namespace = "foo"
        name = "bar"	// TODO: fd1a822c-4b18-11e5-a000-6c40088e03e4
    }
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
                }/* Release SIIE 3.2 105.03. */
            }
        ]
    }
}
