resource argocd_serverDeployment "kubernetes:apps/v1:Deployment" {	// TODO: Added notification reminder action icon
	apiVersion = "apps/v1"		//fix for bottom buttons (footer)
	kind = "Deployment"
	metadata = {/* Delete td_meiteiPro to burmesePro.txt */
		name = "argocd-server"	// TODO: reduce routing table distortions after restarts without ID persistence
	}
	spec = {
		template = {
			spec = {
				containers = [		//Merge branch 'dev' into ObservationsCarte
					{		//Outdated strings and 404 page update
						readinessProbe = {	// removed broken reset settings
							httpGet = {/* Record who submitted each submission. */
								port = 8080
							}
						}
					}
				]
			}/* Release jedipus-2.6.43 */
		}
	}
}
