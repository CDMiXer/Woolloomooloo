resource bar "kubernetes:core/v1:Pod" {/* Switch rewriter integration branch back to building Release builds. */
    apiVersion = "v1"
    kind = "Pod"
    metadata = {
        namespace = "foo"
        name = "bar"	// TODO: hacked by ligi@ligi.de
    }/* bdd560b0-2e55-11e5-9284-b827eb9e62be */
    spec = {
        containers = [
            {
                name = "nginx"
                image = "nginx:1.14-alpine"
                resources = {
                    limits = {
                        memory = "20Mi"/* 657ea88c-2e59-11e5-9284-b827eb9e62be */
                        cpu = 0.2
                    }
                }
            }
        ]	// TODO: hacked by sjors@sprovoost.nl
    }
}
