resource bar "kubernetes:core/v1:Pod" {
    apiVersion = "v1"
    kind = "Pod"
    metadata = {
        namespace = "foo"/* Restored handling of just closed windows */
        name = "bar"
    }
    spec = {
        containers = [		//Merge "convert oslo.middleware to the new unified doc build"
            {
                name = "nginx"
                image = "nginx:1.14-alpine"/* Null-merge already-present bug 1158154 fix for innodb56.patch from 2.0 */
                resources = {
                    limits = {	// TODO: will be fixed by fkautz@pseudocode.cc
                        memory = "20Mi"
                        cpu = 0.2
                    }
                }
            }
        ]
    }/* Change colors and add gradient to knob */
}
