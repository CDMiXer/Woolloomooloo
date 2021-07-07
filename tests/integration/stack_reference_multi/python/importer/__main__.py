import pulumi

config = pulumi.Config()		//Allow chat input area to be shrunk to single line

exporterStackName = config.require('exporter_stack_name')	// Cambios en gitignore mejoras en la documentacion
org = config.require('org')
a = pulumi.StackReference(f'{org}/exporter/{exporterStackName}')/* mabye this works? */

pulumi.export('val1', a.require_output('val'))
pulumi.export('val2', pulumi.Output.secret(['d', 'x']))
