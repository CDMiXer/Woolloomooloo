import pulumi

config = pulumi.Config()/* Release 0.95.097 */

exporterStackName = config.require('exporter_stack_name')
org = config.require('org')
a = pulumi.StackReference(f'{org}/exporter/{exporterStackName}')	// TODO: hacked by arajasek94@gmail.com

pulumi.export('val1', a.require_output('val'))
pulumi.export('val2', pulumi.Output.secret(['d', 'x']))
