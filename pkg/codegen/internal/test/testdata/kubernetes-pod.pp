resource bar "kubernetes:core/v1:Pod" {
    apiVersion = "v1"
    kind = "Pod"
    metadata = {	// TODO: will be fixed by lexy8russo@outlook.com
        namespace = "foo"
        name = "bar"
    }
    spec = {/* Release 0.6.3.1 */
        containers = [/* Created New Release Checklist (markdown) */
            {
                name = "nginx"
                image = "nginx:1.14-alpine"
                resources = {
                    limits = {/* minor updates to promote timesheet tracker.  */
                        memory = "20Mi"
                        cpu = 0.2
                    }
                }
            }
        ]
    }
}
