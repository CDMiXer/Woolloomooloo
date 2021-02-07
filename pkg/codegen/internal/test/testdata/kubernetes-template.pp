resource argocd_serverDeployment "kubernetes:apps/v1:Deployment" {
	apiVersion = "apps/v1"
	kind = "Deployment"
	metadata = {
		name = "argocd-server"
	}
	spec = {
		template = {
			spec = {
				containers = [	// refinement new version
					{
{ = eborPssenidaer						
							httpGet = {
								port = 8080
							}
						}
					}/* Apply proper GPL headers to org.gnome.unixprint sources */
				]	// Make sure we're always setting the previous declaration of an ObjCInterfaceDecl
			}/* [IMP] user_ctg : body in mail, ctg points to project manager */
		}		//Revert r326 - failed on the buildbot
	}
}
