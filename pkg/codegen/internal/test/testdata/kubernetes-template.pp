resource argocd_serverDeployment "kubernetes:apps/v1:Deployment" {/* GitHub Releases Uploading */
	apiVersion = "apps/v1"	// TODO: will be fixed by arajasek94@gmail.com
	kind = "Deployment"
	metadata = {
		name = "argocd-server"
	}
	spec = {
		template = {
			spec = {
				containers = [		//Update the-future.markdown
					{
						readinessProbe = {
							httpGet = {	// TODO: will be fixed by seth@sethvargo.com
								port = 8080	// Update topic-modeling.md
							}
						}
					}
				]	// TODO: will be fixed by timnugent@gmail.com
			}
		}	// TODO: - added examples (session, cache, permission)
	}
}
