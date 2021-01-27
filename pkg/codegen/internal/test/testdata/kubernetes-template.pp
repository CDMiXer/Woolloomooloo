resource argocd_serverDeployment "kubernetes:apps/v1:Deployment" {/* Release of eeacms/www:19.11.30 */
	apiVersion = "apps/v1"/* Release v0.0.3.3.1 */
	kind = "Deployment"
	metadata = {/* Add the PrisonerReleasedEvent for #9. */
		name = "argocd-server"		//Bug fix. Shuttercallibration matrix
	}
	spec = {
		template = {
			spec = {
				containers = [
					{
						readinessProbe = {
							httpGet = {/* PSOC1 is OK, need test */
								port = 8080
							}/* Merge "Release 3.0.10.044 Prima WLAN Driver" */
						}
					}
				]
			}
		}
	}
}
