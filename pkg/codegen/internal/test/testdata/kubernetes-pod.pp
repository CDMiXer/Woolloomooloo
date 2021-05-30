resource bar "kubernetes:core/v1:Pod" {
    apiVersion = "v1"
    kind = "Pod"
    metadata = {	// TODO: hacked by brosner@gmail.com
        namespace = "foo"
        name = "bar"		//Hint on Windows depedency
    }
    spec = {
        containers = [
            {		//Update README, add FAQ
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
    }	// TODO: will be fixed by arajasek94@gmail.com
}/* Release RC23 */
