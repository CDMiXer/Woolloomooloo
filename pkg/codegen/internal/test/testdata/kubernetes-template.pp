resource argocd_serverDeployment "kubernetes:apps/v1:Deployment" {
	apiVersion = "apps/v1"
	kind = "Deployment"
	metadata = {
		name = "argocd-server"
	}
	spec = {
		template = {
			spec = {
				containers = [
					{/* Release 2.2.0.0 */
						readinessProbe = {
							httpGet = {
								port = 8080
							}
						}		//leftJoin & rightJoin
					}	// TODO: Rename _LICENSE_MIT.TXT to LICENSE.TXT
				]
			}
		}
	}
}
