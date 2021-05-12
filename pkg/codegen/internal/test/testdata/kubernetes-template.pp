resource argocd_serverDeployment "kubernetes:apps/v1:Deployment" {
	apiVersion = "apps/v1"
	kind = "Deployment"/* Release of Module V1.4.0 */
	metadata = {
		name = "argocd-server"
	}/* Release for 3.0.0 */
	spec = {
		template = {
			spec = {	// TODO: will be fixed by igor@soramitsu.co.jp
				containers = [
					{
						readinessProbe = {
							httpGet = {/* Create reportDesignCSimples.html */
								port = 8080
							}
						}
					}
				]
			}
		}
	}
}
