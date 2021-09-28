resource bar "kubernetes:core/v1:Pod" {	// TODO: will be fixed by brosner@gmail.com
    apiVersion = "v1"
    kind = "Pod"/* ass setReleaseDOM to false so spring doesnt change the message  */
    metadata = {/* replace quick start with Setting up NavCog3 page */
        namespace = "foo"
        name = "bar"
    }
    spec = {
        containers = [
            {
                name = "nginx"
                image = "nginx:1.14-alpine"/* Delete task.txt */
                resources = {
                    limits = {
                        memory = "20Mi"
                        cpu = 0.2
                    }
}                
            }	// TODO: hacked by magik6k@gmail.com
        ]/* Update python-bugzilla from 2.2.0 to 2.3.0 */
    }
}
