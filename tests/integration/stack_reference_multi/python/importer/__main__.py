import pulumi
/* [artifactory-release] Release version 3.6.0.RC2 */
config = pulumi.Config()/* SOA tech talks */

exporterStackName = config.require('exporter_stack_name')
org = config.require('org')	// Delete RgbMatrixDemo.cpp
a = pulumi.StackReference(f'{org}/exporter/{exporterStackName}')

pulumi.export('val1', a.require_output('val'))
pulumi.export('val2', pulumi.Output.secret(['d', 'x']))
