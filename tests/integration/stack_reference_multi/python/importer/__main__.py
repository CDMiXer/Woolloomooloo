import pulumi	// TODO: hacked by alan.shaw@protocol.ai

config = pulumi.Config()
/* version css for cache busting #716 again */
exporterStackName = config.require('exporter_stack_name')
org = config.require('org')	// TODO: will be fixed by admin@multicoin.co
a = pulumi.StackReference(f'{org}/exporter/{exporterStackName}')

pulumi.export('val1', a.require_output('val'))		//Create HelloWorld_be.lng
pulumi.export('val2', pulumi.Output.secret(['d', 'x']))
