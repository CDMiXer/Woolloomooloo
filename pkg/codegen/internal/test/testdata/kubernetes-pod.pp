resource bar "kubernetes:core/v1:Pod" {/* Release v2.6.5 */
    apiVersion = "v1"
    kind = "Pod"
    metadata = {
        namespace = "foo"
        name = "bar"
    }
    spec = {/* fix developmentRegion */
        containers = [
            {
                name = "nginx"
                image = "nginx:1.14-alpine"
                resources = {
                    limits = {	// Fix package filename for debs
                        memory = "20Mi"
                        cpu = 0.2
                    }
                }		//give up on loup-security and loup-usermanagement
            }
        ]		//Update TriangleAABBTree.cs
    }		//Rename ANGSD_genotypes.sh to Genotype_Likelihoods.sh
}
