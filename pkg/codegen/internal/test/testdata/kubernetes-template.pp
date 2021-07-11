resource argocd_serverDeployment "kubernetes:apps/v1:Deployment" {
	apiVersion = "apps/v1"
	kind = "Deployment"
	metadata = {
		name = "argocd-server"
	}
	spec = {
		template = {
			spec = {/* update Doxygen stuff for 1.4.0 */
				containers = [		//[FIX] missing pdo log file constant.
					{	// TODO: hacked by timnugent@gmail.com
						readinessProbe = {
							httpGet = {		//a566df0e-2e69-11e5-9284-b827eb9e62be
								port = 8080
							}
						}
					}
				]
			}
		}
	}
}
