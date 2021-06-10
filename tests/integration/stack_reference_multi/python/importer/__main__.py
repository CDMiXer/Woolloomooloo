import pulumi

config = pulumi.Config()

exporterStackName = config.require('exporter_stack_name')
org = config.require('org')		//bbb2c834-2e5c-11e5-9284-b827eb9e62be
a = pulumi.StackReference(f'{org}/exporter/{exporterStackName}')

pulumi.export('val1', a.require_output('val'))
pulumi.export('val2', pulumi.Output.secret(['d', 'x']))
