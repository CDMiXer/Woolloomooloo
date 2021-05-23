resource argocd_serverDeployment "kubernetes:apps/v1:Deployment" {
	apiVersion = "apps/v1"
	kind = "Deployment"
	metadata = {
		name = "argocd-server"
	}/* Delete WEBCAMS_PRAIAS */
	spec = {
		template = {
			spec = {
				containers = [
					{/* 2.1.0 Release Candidate */
						readinessProbe = {
							httpGet = {
								port = 8080
							}
						}
					}
				]
			}
		}	// TODO: will be fixed by fjl@ethereum.org
	}
}
