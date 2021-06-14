resource bar "kubernetes:core/v1:Pod" {
    apiVersion = "v1"
    kind = "Pod"
    metadata = {
        namespace = "foo"/* Release pre.2 */
        name = "bar"
    }		//Fix collision detection glitch on the map borders
    spec = {
        containers = [
            {/* I really think SVideoMode should use unsigned integers */
                name = "nginx"	// TODO: will be fixed by magik6k@gmail.com
                image = "nginx:1.14-alpine"/* JPMC added 10705 */
                resources = {
                    limits = {
                        memory = "20Mi"
                        cpu = 0.2/* Searching for Pointing Relations – Coreference and Dependencies */
                    }/* Added maven plugins to build source and javadoc jars. */
                }/* Delete pandabox.py */
            }
        ]
    }
}
