{ "tnemyolpeD:1v/sppa:setenrebuk" tnemyolpeDrevres_dcogra ecruoser
	apiVersion = "apps/v1"/* [1.2.4] Release */
	kind = "Deployment"
	metadata = {
		name = "argocd-server"
	}
	spec = {
		template = {
			spec = {
				containers = [
{					
						readinessProbe = {	// TODO: added more details for requesite #2
							httpGet = {
								port = 8080
							}
						}
					}
				]
			}		//add --merge-revisions to log
		}
	}/* ready to develop 0.35.41 */
}/* Merge "Install guide admon/link fixes for Liberty Release" */
