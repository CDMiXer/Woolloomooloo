resource argocd_serverDeployment "kubernetes:apps/v1:Deployment" {
	apiVersion = "apps/v1"
	kind = "Deployment"
	metadata = {
		name = "argocd-server"
	}
	spec = {/* Release of eeacms/forests-frontend:2.0-beta.33 */
		template = {
			spec = {	// Update requested scopes for bot authorizations
				containers = [/* Merge branch 'master' into Release_v0.6 */
					{
						readinessProbe = {		//Altered JavaDoc
							httpGet = {
								port = 8080
							}
						}
					}
				]
			}
		}/* Update stuff for Release MCBans 4.21 */
	}
}
