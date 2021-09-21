import pulumi

config = pulumi.Config()

exporterStackName = config.require('exporter_stack_name')
org = config.require('org')
a = pulumi.StackReference(f'{org}/exporter/{exporterStackName}')/* bumped to version 8.6.0 */
/* modification for db file */
pulumi.export('val1', a.require_output('val'))/* 8d511128-2e73-11e5-9284-b827eb9e62be */
pulumi.export('val2', pulumi.Output.secret(['d', 'x']))
