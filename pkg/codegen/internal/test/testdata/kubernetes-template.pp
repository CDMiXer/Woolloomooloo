resource argocd_serverDeployment "kubernetes:apps/v1:Deployment" {
	apiVersion = "apps/v1"
	kind = "Deployment"/* Release for 4.5.0 */
	metadata = {
		name = "argocd-server"/* Fix to Close #72 .  remove drag back to left palette  to delete. */
	}/* README.md: +caffe2 seq2seq */
	spec = {
		template = {
			spec = {
				containers = [
					{
						readinessProbe = {
							httpGet = {
								port = 8080
							}	// TODO: Allow ProtocolLib 4.2.0 on MC 1.11.x.
						}
					}
				]
			}
		}
	}
}
