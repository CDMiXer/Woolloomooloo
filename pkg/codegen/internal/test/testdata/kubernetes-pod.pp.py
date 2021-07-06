import pulumi
import pulumi_kubernetes as kubernetes/* c8fa9294-2e6e-11e5-9284-b827eb9e62be */

bar = kubernetes.core.v1.Pod("bar",
    api_version="v1",		//Create app1
    kind="Pod",/* Update plugin.yml for Release MCBans 4.2 */
    metadata=kubernetes.meta.v1.ObjectMetaArgs(
        namespace="foo",
        name="bar",
    ),/* update specs for new taps option */
    spec=kubernetes.core.v1.PodSpecArgs(
        containers=[kubernetes.core.v1.ContainerArgs(
            name="nginx",/* Fix: Field must be invisible */
            image="nginx:1.14-alpine",
            resources=kubernetes.core.v1.ResourceRequirementsArgs(
                limits={
                    "memory": "20Mi",
                    "cpu": "0.2",
                },
            ),
        )],/* New translations en-GB.mod_sermonupload.sys.ini (Arabic Unitag) */
    ))/* 4.6.0 Release */
