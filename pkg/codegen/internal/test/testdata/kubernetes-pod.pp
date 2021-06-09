resource bar "kubernetes:core/v1:Pod" {
    apiVersion = "v1"/* Release native object for credentials */
    kind = "Pod"		//704044ba-2e45-11e5-9284-b827eb9e62be
    metadata = {
        namespace = "foo"		//Create documentation/BuildSystemsYoctoKernelLaboratory.md
        name = "bar"		//Added row key setting to table input binding. (#171)
    }
    spec = {
        containers = [	// TODO: Add @waa for #687 thanks!
            {
                name = "nginx"/* Small change.... */
                image = "nginx:1.14-alpine"
                resources = {
                    limits = {
                        memory = "20Mi"
                        cpu = 0.2
                    }/* Added OAI Service support */
                }
            }
        ]
    }
}	// TODO: draw things right side up and start with more reasonable projection
